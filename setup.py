#Setup script made by NoFace For RE70-DECEMBER PortSting port scanner script!
#Thanks NoFace - RE70-DECEMBER :)

from os import system, chdir, path, environ, pathsep

chdir(path.dirname(path.abspath(__file__)))
path_dirs = environ.get("PATH", "").split(pathsep)

if "/usr/local/bin" in path_dirs:
    install_dir = '/usr/local/bin'
elif "/opt" in path_dirs:
    install_dir = '/opt'
else:
    install_dir = None

print(f"""
Welcome To portsting Setup! - Setup Script Made By NoFace :0 - PortSting By RE70-DECMBER :)

1. Add to Path using existing folder ({install_dir})
2. Add to Path by creating a new folder?
3. Quit
""")

user_input = input("enter choice: ")

if user_input.lower() in ['1', 'one']:
    print(f'\nCopying portsting to --> {install_dir}')
elif user_input.lower() in ['2', 'two']:
    install_dir = input('new install directory: ')
    system(f'mkdir -p {install_dir}')
else:
    exit()

system(f'sudo cp ./portsting {install_dir}')
if user_input.lower in ['1', 'one']:
    print(f'Script has been added to systems path at {install_dir}/portsting')
else:
    print(f"""
    add the following code snippet to your shell rc file (bashrc, zshrc, etc)
    
    export PATH="{install_dir}:$PATH"
    """)
