package visage

import (
   "archive/tar"
   "archive/zip"
   "compress/gzip"
   "encoding/json"
   "fmt"
   "io"
   "os"
   "path/filepath"
   "regexp"
   "strings"
   "testing/fstest"
)

const WingetURL =
   "https://codeload.github.com" +
   "/microsoft/winget-pkgs/tar.gz/refs/heads/master"

var Packages = []Package{
   {
      ID: "Microsoft.VisualC.140.CRT.Headers.Msi",
      Payloads: []Payload{
         // excpt.h
         {Filename: "VC_CRT.Headers.msi"},
      },
   }, {
      ID: "Microsoft.VisualC.140.CRT.x64.Desktop.Msi",
      Payloads: []Payload{
         // libcmt.lib
         {Filename: "VC_CRT.X64.Desktop.msi"},
      },
   }, {
      ID: "Microsoft.VisualC.140.CRT.x64.Store.Msi",
      Payloads: []Payload{
         // msvcrt.lib
         {Filename: "VC_CRT.X64.Store.msi"},
      },
   }, {
      ID: "Microsoft.VisualCpp.Tools.HostX64.TargetX64",
      Payloads: []Payload{
         // nmake.exe
         {Filename: "Microsoft.VisualCpp.Tools.HostX64.TargetX64.vsix"},
      },
   }, {
      ID: "Microsoft.VisualCpp.Tools.HostX64.TargetX64.Resources",
      Payloads: []Payload{
         // clui.dll
         {Filename: "Microsoft.VisualCpp.Tools.HostX64.TargetX64.Resources.enu.vsix"},
      },
   }, {
      ID: "Win10SDK_10.0.19041",
      Payloads: []Payload{
         // ctype.h
         {Filename: "Universal CRT Headers Libraries and Sources-x86_en-us.msi"},
         // security.h
         {Filename: "Windows SDK Desktop Headers x86-x86_en-us.msi"},
         // wldap32.lib
         {Filename: "Windows SDK Desktop Libs x64-x86_en-us.msi"},
         // ws2tcpip.h
         {Filename: "Windows SDK for Windows Store Apps Headers-x86_en-us.msi"},
         // ws2_32.lib
         {Filename: "Windows SDK for Windows Store Apps Libs-x86_en-us.msi"},
         // rc.exe
         {Filename: "Windows SDK for Windows Store Apps Tools-x86_en-us.msi"},
      },
   },
}

var Patterns = []string{
   `Contents\VC\Tools\MSVC\*\bin\Hostx64\x64\nmake.exe`,
   `Program Files\Microsoft Visual Studio *\VC\include\excpt.h`,
   `Program Files\Microsoft Visual Studio *\VC\lib\amd64\msvcrt.lib`,
   `Windows Kits\10\Include\*\shared\winapifamily.h`,
   `Windows Kits\10\Include\*\ucrt\ctype.h`,
   `Windows Kits\10\Include\*\um\WS2tcpip.h`,
   `Windows Kits\10\Lib\*\ucrt\x64\ucrt.lib`,
   `Windows Kits\10\Lib\*\um\x64\WS2_32.Lib`,
   `Windows Kits\10\bin\*\x64\rc.exe`,
}

type Archive struct {
   Strip int
}

func (a Archive) Zip(source, dest string) error {
   rZip, err := zip.OpenReader(source)
   if err != nil {
      return err
   }
   defer rZip.Close()
   for _, entry := range rZip.File {
      if entry.Mode().IsDir() {
         continue
      }
      name := a.strip(dest, entry.Name)
      if name == "" {
         continue
      }
      err := os.MkdirAll(filepath.Dir(name), os.ModeDir)
      if err != nil {
         return err
      }
      open, err := entry.Open()
      if err != nil {
         return err
      }
      file, err := os.Create(name)
      if err != nil {
         return err
      }
      defer file.Close()
      if _, err := file.ReadFrom(open); err != nil {
         return err
      }
   }
   return nil
}

func (a Archive) strip(left, right string) string {
   split := strings.SplitN(right, "/", a.Strip + 1)
   if len(split) <= a.Strip {
      return ""
   }
   return filepath.Join(left, split[a.Strip])
}

type ChannelMan struct {
   ChannelItems []struct {
      Payloads []struct {
         URL string
      }
   }
}

func NewChannelMan(name string) (*ChannelMan, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   chman := new(ChannelMan)
   if err := json.NewDecoder(file).Decode(chman); err != nil {
      return nil, err
   }
   return chman, nil
}

func (c ChannelMan) VisualStudioURL() string {
   return c.ChannelItems[0].Payloads[0].URL
}

type Package struct {
   ID string
   // "", "neutral", "en-US"
   Language string
   Payloads []Payload
}

type Payload struct {
   Filename string
   URL string
}

type VisualStudioMan struct {
   Packages []Package
}

func NewVisualStudioMan(name string) (*VisualStudioMan, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   vsman := new(VisualStudioMan)
   if err := json.NewDecoder(file).Decode(vsman); err != nil {
      return nil, err
   }
   return vsman, nil
}

func (m VisualStudioMan) PayloadURL(pack, payload string) (string, error) {
   for _, vsPackage := range m.Packages {
      if vsPackage.ID != pack {
         continue
      }
      for _, vsPayload := range vsPackage.Payloads {
         if filepath.Base(vsPayload.Filename) == payload {
            return vsPayload.URL, nil
         }
      }
   }
   return "", fmt.Errorf("PayloadURL %v", payload)
}

type WinGet struct {
   fstest.MapFS
}

func NewWinGet(name string) (*WinGet, error) {
   file, err := os.Open(name)
   if err != nil {
      return nil, err
   }
   defer file.Close()
   rGzip, err := gzip.NewReader(file)
   if err != nil {
      return nil, err
   }
   rTar := tar.NewReader(rGzip)
   files := make(fstest.MapFS)
   for {
      cur, err := rTar.Next()
      if err == io.EOF {
         break
      } else if err != nil {
         return nil, err
      }
      if cur.Typeflag != tar.TypeReg {
         continue
      }
      data, err := io.ReadAll(rTar)
      if err != nil {
         return nil, err
      }
      files[cur.Name] = &fstest.MapFile{Data: data}
   }
   return &WinGet{files}, nil
}

func (w WinGet) ChannelURI() (string, error) {
   glob :=
      "winget-pkgs-master/manifests/m/Microsoft/VisualStudio/BuildTools/*/" +
      "Microsoft.VisualStudio.BuildTools.yaml"
   builds, err := w.Glob(glob)
   if err != nil {
      return "", err
   }
   if builds == nil {
      return "", fmt.Errorf("fs.Glob %v", glob)
   }
   latest := w.MapFS[builds[len(builds) - 1]].Data
   re := regexp.MustCompile(`--channelUri (\S+)`)
   channelUri := re.FindSubmatch(latest)
   if channelUri == nil {
      return "", fmt.Errorf("FindSubmatch %v", re)
   }
   return string(channelUri[1]), nil
}
