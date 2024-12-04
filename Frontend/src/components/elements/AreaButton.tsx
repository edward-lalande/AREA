import React from "react";
import { Button, ButtonProps } from "@mui/material";
import GoogleIcon from "@mui/icons-material/Google";
import LogoutIcon from '@mui/icons-material/Logout';

interface AreaButtonProps extends ButtonProps {
    text: string;
}

const AreaButton: React.FC<AreaButtonProps> = ({ text, ...props }) => {
    return (
        <Button
        variant="contained"
        fullWidth
        sx={{
                backgroundColor: "#000",
                color: "#fff",
                borderRadius: 5,
                fontWeight: "bold",
                py: 1.5,
                fontSize: "1rem",
                textTransform: "none",
                maxWidth: 400,
                ...props.sx
            }}
            {...props}
        >
            {text}
        </Button>
    );
};

const CreateButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
        variant="contained"
        fullWidth
        href="/"
        sx={{
                backgroundColor: "#000",
                color: "#fff",
                borderRadius: 10,
                ml: 1,
                fontWeight: "bold",
                fontSize: "1.2em",
                textTransform: "none",
                width: "10%",
                minWidth: 100,
                height: "65%",
                ...props.sx
            }}
            {...props}
        >
            Create
        </Button>
    )
};

const GoogleButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            startIcon={<GoogleIcon />}
            sx={{
                borderColor: "#888",
                color: "#000",
                borderRadius: 5,
                py: 1.5,
                textTransform: "none",
                fontWeight: "bold",
                fontSize: "1rem",
                mb: 1,
                maxWidth: 400,
                ...props.sx
            }}
            {...props}
        >
            Continue with Google
        </Button>
    );
};

const Logout: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            startIcon={<LogoutIcon />}
            sx={{
                borderColor: "red",
                color: "red",
                borderRadius: 10,
                py: 1.5,
                textTransform: "none",
                fontWeight: "bold",
                fontSize: "1rem",
                mb: 1,
                maxWidth: 150,
                ...props.sx
            }}
            {...props}
        >
            Logout
        </Button>
    );
};
  
export { AreaButton, CreateButton, GoogleButton, Logout };