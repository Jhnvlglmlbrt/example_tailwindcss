# techonology stack:

go, tailwindcss, air

## Installation

```
cd web && npm init -Y
```

```
npm install -D tailwindcss
npx tailwindcss init
```

change in tailwind.config.js 

```
content: ["./views/**/*.{html,js}"],
```
cd web && mkdir styles
touch inputs.css

add in input.css 
```
@tailwind base;
@tailwind components;
@tailwind utilities;
```

then

```
go install github.com/cosmtrek/air@latest
go get -u github.com/cosmtrek/air
air init
```

use this in one terminal

```
npx tailwindcss -i ./styles/input.css -o ./styles/output.css --watch
```

and this in another 

```
air
```