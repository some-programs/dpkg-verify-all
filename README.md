# dpkg-verify-all

A quick hack to verify hashes of all installed dpkg packages and report which
ones doesnt verify. I found a broken header file on my harddrive today and
wrote this to be able to quickly reverify all packages to identify if my /
drive maybe is beginning to fail.


## Example run

Intentionally without sudo to generate more output, also the vagrant package
seems to have installed a broken checksum file:


```sh
$ time go run pkg-verify.go 
Verifying 3299 packages using 8 workers...
Started verification of 100 packagages...
Started verification of 200 packagages...
Started verification of 300 packagages...
Started verification of 400 packagages...
Started verification of 500 packagages...
Started verification of 600 packagages...
Started verification of 700 packagages...
Started verification of 800 packagages...
Started verification of 900 packagages...
Started verification of 1000 packagages...
Started verification of 1100 packagages...
Started verification of 1200 packagages...
Started verification of 1300 packagages...
Started verification of 1400 packagages...
Started verification of 1500 packagages...
Started verification of 1600 packagages...
Started verification of 1700 packagages...
Started verification of 1800 packagages...
Started verification of 1900 packagages...
Started verification of 2000 packagages...
Started verification of 2100 packagages...
Started verification of 2200 packagages...
Started verification of 2300 packagages...
Started verification of 2400 packagages...
Started verification of 2500 packagages...
Started verification of 2600 packagages...
Started verification of 2700 packagages...
Started verification of 2800 packagages...
Started verification of 2900 packagages...
Started verification of 3000 packagages...
Started verification of 3100 packagages...
Started verification of 3200 packagages...

17 package verifications returned non empty output
ca-certificates-java: dpkg: warning: ca-certificates-java: unable to open /etc/default/cacerts for hash: Permission denied
ca-certificates-java: ??5?????? c /etc/default/cacerts
cups-filters: dpkg: warning: cups-filters: unable to open /usr/lib/cups/backend/serial for hash: Permission denied
cups-filters: ??5??????   /usr/lib/cups/backend/serial
fwupd: dpkg: warning: fwupd: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/fwupd.pkla for hash: Permission denied
fwupd: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/fwupd.pkla
google-chrome-stable: dpkg: warning: google-chrome-stable: unable to open /usr/share/doc/google-chrome-stable/changelog.gz for hash: Permission denied
google-chrome-stable: ??5??????   /usr/share/doc/google-chrome-stable/changelog.gz
indicator-sound: dpkg: warning: indicator-sound: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/50-com.canonical.indicator.sound.AccountsService.pkla for hash: Permission denied
indicator-sound: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/50-com.canonical.indicator.sound.AccountsService.pkla
linux-image-4.2.0-17-generic: dpkg: warning: linux-image-4.2.0-17-generic: unable to open /boot/System.map-4.2.0-17-generic for hash: Permission denied
linux-image-4.2.0-17-generic: ??5??????   /boot/System.map-4.2.0-17-generic
linux-image-4.2.0-17-generic: dpkg: warning: linux-image-4.2.0-17-generic: unable to open /boot/vmlinuz-4.2.0-17-generic for hash: Permission denied
linux-image-4.2.0-17-generic: ??5??????   /boot/vmlinuz-4.2.0-17-generic
linux-image-4.4.0-21-generic: dpkg: warning: linux-image-4.4.0-21-generic: unable to open /boot/System.map-4.4.0-21-generic for hash: Permission denied
linux-image-4.4.0-21-generic: ??5??????   /boot/System.map-4.4.0-21-generic
linux-image-4.4.0-21-generic: dpkg: warning: linux-image-4.4.0-21-generic: unable to open /boot/vmlinuz-4.4.0-21-generic for hash: Permission denied
linux-image-4.4.0-21-generic: ??5??????   /boot/vmlinuz-4.4.0-21-generic
linux-image-4.4.0-22-generic: dpkg: warning: linux-image-4.4.0-22-generic: unable to open /boot/System.map-4.4.0-22-generic for hash: Permission denied
linux-image-4.4.0-22-generic: ??5??????   /boot/System.map-4.4.0-22-generic
linux-image-4.4.0-22-generic: dpkg: warning: linux-image-4.4.0-22-generic: unable to open /boot/vmlinuz-4.4.0-22-generic for hash: Permission denied
linux-image-4.4.0-22-generic: ??5??????   /boot/vmlinuz-4.4.0-22-generic
mpd: dpkg: warning: mpd: unable to open /etc/mpd.conf for hash: Permission denied
mpd: ??5?????? c /etc/mpd.conf
network-manager: dpkg: warning: network-manager: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/org.freedesktop.NetworkManager.pkla for hash: Permission denied
network-manager: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/org.freedesktop.NetworkManager.pkla
policykit-desktop-privileges: dpkg: warning: policykit-desktop-privileges: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/com.ubuntu.desktop.pkla for hash: Permission denied
policykit-desktop-privileges: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/com.ubuntu.desktop.pkla
ppp: dpkg: warning: ppp: unable to open /etc/chatscripts/gprs for hash: Permission denied
ppp: ??5?????? c /etc/chatscripts/gprs
ppp: dpkg: warning: ppp: unable to open /etc/chatscripts/pap for hash: Permission denied
ppp: ??5?????? c /etc/chatscripts/pap
sudo: dpkg: warning: sudo: unable to open /etc/sudoers for hash: Permission denied
sudo: ??5?????? c /etc/sudoers
sudo: dpkg: warning: sudo: unable to open /etc/sudoers.d/README for hash: Permission denied
sudo: ??5?????? c /etc/sudoers.d/README
unar: ??5?????? c /etc/bash_completion.d/unar
unity-greeter: dpkg: warning: unity-greeter: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/unity-greeter.pkla for hash: Permission denied
unity-greeter: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/unity-greeter.pkla
unity-webapps-common: dpkg: warning: unity-webapps-common: unable to open /var/lib/polkit-1/localauthority/10-vendor.d/com.canonical.unity.webapps.pkla for hash: Permission denied
unity-webapps-common: ??5??????   /var/lib/polkit-1/localauthority/10-vendor.d/com.canonical.unity.webapps.pkla
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxHeadless for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxHeadless
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxSDL for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxSDL
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxNetNAT for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxNetNAT
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VirtualBox for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VirtualBox
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxNetDHCP for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxNetDHCP
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxVolInfo for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxVolInfo
virtualbox-5.0: dpkg: warning: virtualbox-5.0: unable to open /usr/lib/virtualbox/VBoxNetAdpCtl for hash: Permission denied
virtualbox-5.0: ??5??????   /usr/lib/virtualbox/VBoxNetAdpCtl

 dpkg --verify finished with non zero exit 1 times:
vagrant: dpkg: error: control file 'md5sums' missing value separator
exit status 1

real	1m17.692s
user	4m15.468s
sys	0m17.076s

```
