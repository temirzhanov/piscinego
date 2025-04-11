#1 /bin/bash

find . -name '*.sh' | rev | cut -d '/' -f1 | rev | cut -d '.' -f1  