import React from "react";
import { Box, Button, ButtonProps, Card, CardProps, Typography } from "@mui/material";
import GoogleIcon from "@mui/icons-material/Google";
import LogoutIcon from '@mui/icons-material/Logout';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import { AreaTypography } from "./AreaTypography";
import { AreaBox } from "./AreaBox";

interface Service {
	id: number;
	name: string;
	color: string;
	description: string;
}

interface AreaButtonProps extends ButtonProps {
    text: string;
}

interface ServiceButtonProps extends ButtonProps {
    text: string;
    backgroundColor: string;
}

interface ServiceCardCreateProps extends CardProps {
    text: string;
    backgroundColor: string;
}

interface ServiceCardProps {
    service: Service;
    onClick: (value: Service) => void;
}

const AreaButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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

const AccountButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
    <Button
    variant="contained"
    fullWidth
    sx={{
            backgroundColor: "#000",
            color: "#fff",
            borderRadius: 5,
            width: "15vw",
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

const ServiceCard: React.FC<ServiceCardProps> = ({ service, onClick }) => (
    <Card variant="outlined" sx={{ width: "22vw", height: "18vh", display: "flex", flexDirection: "row", borderRadius: 5 }}>
        <AreaBox sx={{ width: "8vw", height: "18vh" }}>
            <img src={service.name.toLowerCase().replace(' ', '_') + ".png"} width={80} height={80} />
        </AreaBox>
        <Box sx={{ width: "14vw", height: "18vh", backgroundColor: service.color, padding: 2, gap: 2 }}>
            <Typography color="white" sx={{ fontSize: 24, mb: 0.5 }}>
                {service.name}
            </Typography>
            <Typography color="white" sx={{ height: "5.8vh", fontSize: 12 }}>
                {service.description.length > 80 ? service.description.slice(0, 80) + "..." : service.description}
            </Typography>
            <Button
                variant="contained"
                onClick={() => onClick(service)}
                sx={{
                    backgroundColor: "white",
                    color: "black",
                    textTransform: "none",
                    borderRadius: 3,
                    mt: 1,
                }}
            >
                See details
            </Button>
        </Box>
    </Card>
);

const CreateButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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

const AIButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
    <Button
    variant="contained"
    fullWidth
    href="/create-ai"
    sx={{
            backgroundColor: "#352f5f",
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

const AddButton: React.FC<ButtonProps> = (props) => (
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

const Logout: React.FC<ButtonProps> = (props) => (
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

const HelpButton: React.FC<ButtonProps> = (props) => (
    <Button
        variant="outlined"
        fullWidth
        startIcon={<HelpOutlineIcon />}
        sx={{
            borderColor: "black",
            color: "black",
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
            Help
    </Button>
);

const DiscordButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <AreaTypography variant="h6" text={text} sx={{ ml: 2 }}/>
    </Button>
);

const SpotifyButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <AreaTypography variant="h6" text={text} sx={{ ml: 2 }}/>
    </Button>
);

const GithubButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <AreaTypography variant="h6" text={text} sx={{ ml: 2 }}/>
    </Button>
);

const GitlabButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <AreaTypography variant="h6" text={text} sx={{ ml: 3 }}/>
    </Button>
);

const GoogleButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <AreaTypography variant="h6" text={text} sx={{ ml: 3 }}/>
    </Button>
);

const DropboxButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
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
        <img src="dropbox-logo.png" alt="dropbox logo" width={44} height={44}/>
        <AreaTypography variant="h6" text={text} sx={{ ml: 3 }}/>
    </Button>
)

const AsanaButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
    <Button
        variant="outlined"
        fullWidth
        sx={{
            backgroundColor: "#fff",
            color: "black",
            borderRadius: 5,
            border: "3px solid pink",
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
        <img src="asana-logo.png" alt="asana logo" width={70} height={44}/>
        <AreaTypography variant="h6" text={text} sx={{ ml: 3 }}/>
    </Button>
)

const MiroButton: React.FC<AreaButtonProps> = ({ text, ...props }) => (
    <Button
        variant="outlined"
        fullWidth
        sx={{
            backgroundColor: "#faca00",
            color: "black",
            borderRadius: 5,
            border: "3px solid #faca00",
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
        <img src="miro-logo.png" alt="miro logo" width={44} height={44}/>
        <AreaTypography variant="h6" text={text} sx={{ ml: 3 }}/>
    </Button>
)

export { ServiceCard, AIButton, MiroButton, HelpButton, AsanaButton, DropboxButton, AreaButton, DiscordButton, GoogleButton, SpotifyButton, GithubButton, GitlabButton, Logout, CreateButton, AddButton, ServiceButton, AccountButton };
