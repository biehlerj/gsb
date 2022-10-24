# TODOS

As features are implemented they will get their own documentation

## Features

1. Backup keybindings:
   1. WM specific
   2. Non-WM specific
   3. Power related buttons
2. Other desktop related settings:
   1. Background
   2. Peripherals
   3. Notifications
   4. Interface?
   5. Default Applications
   6. Privacy
   7. WM Preferences
   8. Search Providers?

## Equivalent `gsettings` commands for exporting

All export commands use `gsettings list-recursively [setting]`

- Keybindings:
  - `org.gnome.shell.keybindings`
  - `org.gnome.mutter.keybindings`
  - `org.gnome.desktop.wm.keybindings`
  - `org.gnome.settings-daemon.peripherals.keyboard`
  - `org.gnome.settings-daemon.plugins.media-keys`
- Settings:
  - `org.gnome.desktop.background`
  - `org.gnome.desktop.peripherals`
  - `org.gnome.desktop.notifications`
  - `org.gnome.desktop.default-applications`
  - `org.gnome.desktop.privacy`
  - `org.gnome.desktop.wm.preferences`
