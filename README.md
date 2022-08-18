
### This is a basic utility to toggle the DLSS indicator overlay as described in [this arcticle](https://www.pcgamer.com/nvidia-dlss-indicator/)


#### How to build:
- Embed the manifest inside your binary 

      go get github.com/akavel/rsrc
      rsrc -manifest DLSS_indicator_toggle.manifest -o rsrc.syso
      
- Build your binary 

        go build -ldflags="-H windowsgui"
        
        

##### Copyright (C) 2022 [Watn3y](https://github.com/Watn3y/)
###### Thanks to Aniums


##### Licensed under [The Unlicense](LICENSE)

