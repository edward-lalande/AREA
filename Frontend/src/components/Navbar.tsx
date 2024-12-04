import { Avatar, IconButton } from "@mui/material";
import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { CreateButton } from "./elements/AreaButton";

const Navbar: React.FC = () => {

	return (

        <AreaBox sx={{ height: "8vh", width: "100%", flexDirection: "row" , mt: 1 }}>

            <AreaBox sx={{ height: "100%", width: "20vw", alignItems: "start" }}>
                <AreaLink href="/" variant="h4" text="AREA" sx={{ ml: 4, textDecoration: "none" }} />
            </AreaBox>

            <AreaBox sx={{ height: "100%", width: "80vw", flexDirection: "row", gap: 4, justifyContent: "end", pr: 3 }}>

                <AreaLink href="/" text="Services" sx={{ fontSize: "1.5em", textDecoration: "none" }} />
                <AreaLink href="/" text="Areas" sx={{ fontSize: "1.5em", textDecoration: "none" }} /> 
                <CreateButton text="Create" />

                <IconButton href="/account">
                    <Avatar sx={{ border: "3px solid black", width: 48, height: 48 }}>P</Avatar>
                </IconButton>

            </AreaBox>

        </AreaBox>

	);

};

export default Navbar;