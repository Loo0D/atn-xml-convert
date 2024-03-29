# atn-xml-convert

Convert TOML ballistic settings file to `profiles.xml` for ATN X-Sight 4K Scope.

## Update

1) Convert profiles.xml file from the scope into a TOML file
2) The new version changed flags to `-t` for input TOML file and added `-x` for input XML file
3) You may now omit "unknown" settings in your TOML file, and they will be set to default values so your TOML file can be a bit cleaner:

Old:
```toml
[one]
  profile_name = "F: Hornady 180 SST"
  drag_function = 1
  bc = 0.48
  bullet_weight = 180.0
  init_velocity = 2443.0
  sight_height = 2.5
  zeroing_distance = 70.5
  reticle_offset_x = -54
  reticle_offset_y = 229
  msp_acc_border_value = 5000               # Don't know what 
  acc_border_up_cross_counter_min = 4       # these do  
  msp_acc_border_up_cross_counter_max = 30  # but they don't
  msp_acc_border_down_cross_counter_min = 1 # seem to change
```

New:
```toml
[one]
profile_name = "F: Hornady 180 SST"
drag_function = 1
bc = 0.48
bullet_weight = 180.0
init_velocity = 2443.0
sight_height = 2.5
zeroing_distance = 70.5
reticle_offset_x = -54
reticle_offset_y = 229
```

See "Unknown Settings" below for more info.

## Why

Manually editing `profiles.xml` files exported by ATN X-Sight 4K Scope is [painful](https://atnowners.com/thread/3520/atn-sight-profiles-editing-guideline).

Since ATN only allows you to keep 6 profiles loaded in the scope at a time, if you reload for several calibers, you soon end up having to manage several profile files. Keeping them updated by importing them into the scope and editing there via the menu is a hassle.

But what of the Obsidian app from ATN you say? The app has a serious problem - it aggressively rounds up values you enter. For example, a flathead air rifle pellet has a BC of 0.0250. Entered in the app, it rounds it up to 0.1 and you could get a very wrong ballistic calculation.

A solution I wanted was a simple text format for my ballistic settings that I could keep together with my reloading data, check into version control and be able to quickly get it uploaded to the scope. 

## What

See `example-settings.toml` for a TOML template. Amend it to suit your ballistic profiles. It's set up for 6, the maximum allowed by the scope - if you have fewer just call them "Unused 1", "Unused 2" etc.

All settings in the XML file are metric (except grains), but if you want to work in yards/fps/inches set the flags at the top of the TOML file to `true`. They apply to the **entire file**, so all the 6 profiles you will need to set in inches or fps etc. If flags are `false` then specify values in native XML format.

Alternatively, export your existing profiles from the scope into `profiles.xml` and use this tool to convert it to a TOML file for easy editing and archival.

## Usage: TOML to XML

Once your TOML file is ready, run the binary and tell it where you TOML file is:

```shell
./atn-xml-convert -t myBallisticSettings.toml # produces a file called profiles.xml in current directory, ready to be imported into the scope
```

or

```shell
./atn-xml-convert -t myBallisticSettings.toml -o ~/Desktop/profiles.xml
```

The binary will write `profiles.xml` in its directory (or where you specify with `-o` flag) that you can put on the microSD card and import into the scope. The file will need to be named `profiles.xml` for the scope to recognise it.

## Usage: XML to TOML

If you want to export your profiles.xml file to a TOML file, run the binary and tell it where you XML file is:

```shell
./atn-xml-convert -x profiles.xml # produces a file called profiles.toml in current directory
```

or

```shell
./atn-xml-convert -x profiles.xml -o ~/Desktop/mySettings.toml
```

**NOTE** : by default, the TOML file will have units as they are in the "raw" profiles.xml file, which is to say metric. If you want your converted TOML file to have imperial units, set the appropriate flags: `-yards` (zero distance) and/or `-fps` (muzzle velocity) and/or `-inches` (sight height) as follows:

```shell
./atn-xml-convert -x profiles.xml -o ~/Desktop/profiles.toml -yards -fps -inches # everything is in imperial units
```

```shell
./atn-xml-convert -x profiles.xml -o ~/Desktop/profiles.toml -fps # everything is metric except muzzle velocity
```

## Unknown Settings

I don't know what these do, they don't look ballistics-related:

- `msp_acc_border_value=5000`
- `acc_border_up_cross_counter_min=4`
- `msp_acc_border_up_cross_counter_max=30`
- `msp_acc_border_down_cross_counter_min=1`

They appear to be constant, i.e. don't change with other settings, so it's probably best to leave them as is.

## Errors/Contributions

There isn't a lot here to go wrong, but if you find errors notify me or submit a pull request.

Additional features: instead of wasting effort on making a web app or a GUI for this, we should probably use the time to tell ATN to 

1) not use an undocumented and obfuscated semi-proprietary format for its settings exports (yeah, it could have been worse)
2) get its' Obsidian rounding errors fixed
