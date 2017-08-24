installPackages (){
    curl http://s.sudre.free.fr/Software/files/Packages.dmg -o packages.dmg
    sudo hdiutil attach ./packages.dmg
    sudo installer -package "/Volumes/Packages 1.1.3/packages/Packages.pkg" -target /
    sudo hdiutil detach "/Volumes/Packages 1.1.3"
}

installPackages
