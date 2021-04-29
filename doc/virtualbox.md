# VirtualBox

## Create Virtual Machine

Set Base memory to 4096 MB. Set video memory to 128 MB.

## Guest additions Windows

1. Optical Disk Selector, Add
2. Devices, insert guest additions cd image
3. Install guest additions
4. Devices, optical drives, remove disk from virtual drive
5. Devices, shared folders, shared folders settings
6. Add new shared folder
7. Folder path `D:\virtualbox shared`
8. Make permanent, true
9. Create shortcut `\\vboxsvr\virtualbox shared`
10. Snapshot

## Preferences

1. File
2. Preferences
3. Default machine folder `D:\VirtualBox VMs`

## References

- <https://superuser.com/questions/1108085>
- <https://virtualbox.org>
