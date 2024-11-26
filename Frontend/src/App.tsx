import Home from './components/Home';
import LoginForm from './components/LoginForm';
import SignupForm from './components/SignupForm';
import defaultTheme from './themes/defaultTheme';
import ThemeProvider from './components/ThemeProvider';

import "./App.css"
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const App: React.FC = () => {

    const router = createBrowserRouter([
        {
            path: "/",
            element: <Home />
        },
        {
            path: "/login",
            element: <LoginForm />
        },
        {
            path: "/signup",
            element: <SignupForm />
        }
    ]);

    return (
        <ThemeProvider theme={defaultTheme}>
            <RouterProvider router={router} />
        </ThemeProvider>
    );
};

export default App;
