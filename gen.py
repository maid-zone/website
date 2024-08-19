# small script to generate ansi stuffs
from rich.console import Console

console = Console(highlight=False)
with console.capture() as capt:
    console.print("""
                             o8o        .o8                                                 
                             `"'       "888                                                 
ooo. .oo.  .oo.    .oooo.   oooo   .oooo888        oooooooo  .ooooo.  ooo. .oo.    .ooooo.  
`888P"Y88bP"Y88b  `P  )88b  `888  d88' `888       d'""7d8P  d88' `88b `888P"Y88b  d88' `88b 
 888   888   888   .oP"888   888  888   888         .d8P'   888   888  888   888  888ooo888 
 888   888   888  d8(  888   888  888   888  .o.  .d8P'  .P 888   888  888   888  888    .o 
o888o o888o o888o `Y888""8o o888o `Y8bod88P" Y8P d8888888P  `Y8bod8P' o888o o888o `Y8bod8P' 
                                                                                            
small collective of friends

who is here:
- [#ff4151][link=https://laptopc.at]laptopcat[/link][/#ff4151]
    owner
- [#ff4151][link=https://delta.maid.zone]deltaost[/link][/#ff4151]
    member
""")

with open("out.ans", "w") as file:
    file.write(capt.get())
