import React from "react";
import { Button, ButtonProps } from "@mui/material";
import GoogleIcon from "@mui/icons-material/Google";
import LogoutIcon from '@mui/icons-material/Logout';
import { AreaTypography } from "./AreaTypography";

interface AreaButtonProps extends ButtonProps {
    text: string;
}

interface ServiceButtonProps extends ButtonProps {
    text: string;
    backgroundColor: string;
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

const ServiceButton: React.FC<ServiceButtonProps> = ({ text, backgroundColor, ...props }) => {
    return (
        <Button
        variant="contained"
        fullWidth
        sx={{
                backgroundColor,
                color: "#fff",
                borderRadius: 8,
                fontWeight: "bold",
                height: "8vw",
                width: "8vw",
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

const CreateButton: React.FC<AreaButtonProps> = ({ text, ...props }) => {
    return (
        <Button
        variant="contained"
        fullWidth
        href="/create"
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
            {text}
        </Button>
    )
};

const AddButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
        variant="contained"
        fullWidth
        sx={{
            color: "black",
            backgroundColor: "white",
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
            Add
        </Button>
    )
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
                maxWidth: 150,
                py: 1.5,
                textTransform: "none",
                fontWeight: "bold",
                fontSize: "1rem",
                mb: 1,
                ...props.sx
            }}
            {...props}
            >
                Logout
        </Button>
    );
};

const DiscordButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                borderColor: "#fff",
                backgroundColor: "#5865f2",
                color: "#fff",
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
            <img src="discord-logo.png" alt="discord logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Discord" sx={{ ml: 2 }}/>
        </Button>
    );
};

const SpotifyButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                borderColor: "#1db954",
                backgroundColor: "#1db954",
                color: "#fff",
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
            <img src="spotify-logo.png" alt="spotify logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Spotify" sx={{ ml: 2 }}/>
        </Button>
    );
};

const GithubButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                borderColor: "#fff",
                backgroundColor: "#000",
                color: "#fff",
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
            <img src="logo-github.png" alt="github logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Github" sx={{ ml: 2 }}/>
        </Button>
    );
};

const GitlabButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                backgroundColor: "#fff",
                color: "#FC6D27",
                borderRadius: 5,
                border: "3px solid #FC6D27",
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
            <img src="logo-gitlab.png" alt="gitlab logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Gitlab" sx={{ ml: 3 }}/>
        </Button>
    );
};

const GoogleButton: React.FC<ButtonProps> = (props) => {
    return (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                backgroundColor: "#fff",
                color: "black",
                borderRadius: 5,
                border: "3px solid black",
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
            <img src="google-logo.png" alt="google logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Google" sx={{ ml: 3 }}/>
        </Button>
    );
};

const DropboxButton: React.FC<ButtonProps> = (props) => (
        <Button
            variant="outlined"
            fullWidth
            sx={{
                backgroundColor: "#fff",
                color: "black",
                borderRadius: 5,
                border: "3px solid blue",
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
            <img src="dropbox-logo.png" alt="google logo" width={44} height={44}/>
            <AreaTypography variant="h6" text="Continue with Dropbox" sx={{ ml: 3 }}/>
        </Button>
)

export { DropboxButton, AreaButton, DiscordButton, GoogleButton, SpotifyButton, GithubButton, GitlabButton, Logout, CreateButton, AddButton, ServiceButton };
