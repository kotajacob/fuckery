# fuckery - General unicode tomfoolery.

Fuckery is a go library and cli tool for applying those 𝓌ℯ𝒾𝓇𝒹 𝕦𝕟𝕚𝕔𝕠𝕕𝕖 𝒔𝒕𝒚𝒍𝒆𝒔 to
text.

## Usage

Fuckery reads from stdin and prints to stdout.

fuckery [style]

styles:
 * Strike
 * Underline
 * BoldSans
 * BoldSerif
 * ItalicSans
 * ItalicSerif
 * BoldItalicSans
 * BoldItalicSerif
 * Double
 * Cursive
 * Fraktur

## Build and Install

Build dependencies:  
 * golang
 * make
 * sed
 * scdoc

Optionally configure `config.mk` to specify a different install location.  
Defaults to `/usr/local/`

`sudo make all`

## Uninstall

`sudo make uninstall`
