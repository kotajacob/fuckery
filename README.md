# fuckery - General unicode tomfoolery.

[![godocs.io](https://godocs.io/git.sr.ht/~kota/fuckery?status.svg)](https://godocs.io/git.sr.ht/~kota/fuckery)

Fuckery is a go library and cli tool for applying those 𝓌ℯ𝒾𝓇𝒹 𝕦𝕟𝕚𝕔𝕠𝕕𝕖 𝒔𝒕𝒚𝒍𝒆𝒔 to
text. Your terminal's stdout may trim some of the styles, try printing to a file
if you have issues.

## Usage

Fuckery reads from stdin and prints to stdout.

fuckery [style]

styles:
 * S̶t̶r̶i̶k̶e̶
 * U̲n̲d̲e̲r̲l̲i̲n̲e̲
 * 𝗕𝗼𝗹𝗱𝗦𝗮𝗻𝘀
 * 𝐁𝐨𝐥𝐝𝐒𝐞𝐫𝐢𝐟
 * 𝘐𝘵𝘢𝘭𝘪𝘤𝘚𝘢𝘯𝘴
 * 𝐼𝑡𝑎𝑙𝑖𝑐𝑆𝑒𝑟𝑖𝑓
 * 𝘽𝙤𝙡𝙙𝙄𝙩𝙖𝙡𝙞𝙘𝙎𝙖𝙣𝙨
 * 𝑩𝒐𝒍𝒅𝑰𝒕𝒂𝒍𝒊𝒄𝑺𝒆𝒓𝒊𝒇
 * 𝔻𝕠𝕦𝕓𝕝𝕖
 * 𝒞𝓊𝓇𝓈𝒾𝓋ℯ
 * 𝔉𝔯𝔞𝔨𝔱𝔲𝔯
 * Z̸̜͇̑̂̓͂a̵̴̜̽ͬͫͅl͚͍̔̀̂ͫ͜g̩ͮ̈͐̒́͏ò̷̗̊͛̀͛ ̲̣̔̈́̅̒͢7͍̣̹̝̌̏͞ (Zalgo 7 -- greater numbers have more zalg)

## Build and Install

Build dependencies:  
 * golang
 * make
 * sed
 * scdoc

Optionally configure `config.mk` to specify a different install location.  
Defaults to `/usr/local/`

`sudo make install`

## Uninstall

`sudo make uninstall`
