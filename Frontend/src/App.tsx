import Home from "./components/Home";
import Account from "./components/Account";
import LoginForm from "./components/LoginForm";
import SignupForm from "./components/SignupForm";
import defaultTheme from "./themes/defaultTheme";
import ThemeProvider from "./components/ThemeProvider";

import "./App.css"
import { createBrowserRouter, RouterProvider, Navigate } from "react-router-dom";
import { ProtectedRoute } from "./components/PrivateRoute";

const App: React.FC = () => {

    const router = createBrowserRouter([
        {
            path: "/",
            element: <ProtectedRoute children={<Home/>} />
        },
        {
            path: "/login",
            element: <LoginForm />
        },
        {
            path: "/signup",
            element: <SignupForm />
        },
        {
            path: "/account",
            element: <ProtectedRoute children={<Account/>} />
        },
        {
            path: "*",
            element: <Navigate to="/" replace />
        }
    ]);

    return (
        <ThemeProvider theme={defaultTheme}>
            <RouterProvider router={router} />
        </ThemeProvider>
    );
};

export default App;
