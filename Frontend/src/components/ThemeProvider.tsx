import defaultTheme from '../themes/defaultTheme';
import CssBaseline from '@mui/material/CssBaseline';
import { ReactNode } from 'react';
import { ThemeProvider as MUIThemeProvider } from '@mui/material/styles';

interface ThemeProviderProps {
    children: ReactNode;
    theme?: typeof defaultTheme;
}

const ThemeProvider = ({ children, theme = defaultTheme }: ThemeProviderProps) => {
    return (
        <MUIThemeProvider theme={theme}>
            <CssBaseline />
            {children}
        </MUIThemeProvider>
    );
};

export default ThemeProvider;
