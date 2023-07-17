# Theming

## Theme provider

Use ThemeProvider component in order to inject a theme into your application.
ThemeProvider relies on the context feature of React to pass the theme down to the components, so you need to make sure that ThemeProvider is a parent of the components you are trying to customize.

### Them configuration variables

.palette
.typography
.spacing
.breakpoints
.zIndex
.transitions
.components

### Custom variables

```ts
declare module '@mui/material/styles' {
    interface Theme {
        status: {
            danger: string;
        };
    }
    // allow configuration using `createTheme`
    interface ThemeOptions {
        status?: {
            danger?: string;
        }
    }
}
```

### Theme builder

The community has built great tools to build a theme

### Accessing the theme in a component

You can access the theme variables inside your function React components using the useTheme hook

```ts
import {useTheme} from '@mui/material/styles';

function DeepChild() {
    const theme = useTheme();
    return <span>{`spacing ${theme.spacing}`}</span>
}
```

### Nesting the theme

You can nest multiple theme providers

### API

createTheme(options, ...args) => theme
Generate a theme base on the options received. Then, pass it as a prop to ThemeProvider



