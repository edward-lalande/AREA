import React from "react";
import { Button, ButtonProps } from "@mui/material";
import GoogleIcon from "@mui/icons-material/Google";

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
  
export { AreaButton, GoogleButton };
