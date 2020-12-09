use toml::ser;

fn main() -> Result<(), ser::Error> {
   let m = toml::toml!{
      [[Winter]]
      December = [
         { Sunday = 6 },
         { Monday = 7 }
      ]
      January = [
         { Sunday = 6 },
         { Monday = 7 }
      ]
      [[Spring]]
      March = [
         { Sunday = 6 },
         { Monday = 7 }
      ]
      April = [
         { Sunday = 6 },
         { Monday = 7 }
      ]
   };
   
   let s = toml::to_string(&m)?;
   //let s = toml::to_string_pretty(&m)?;
   
   println!("{}", s);
   Ok(())
}
