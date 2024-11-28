import { Avatar } from "@mui/material";
import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { CreateButton } from "./elements/AreaButton";

const Navbar: React.FC = () => {

	return (

        <AreaBox sx={{ height: "8vh", width: "100%", flexDirection: "row" }}>

            <AreaBox sx={{ height: "100%", width: "20vw", alignItems: "start" }}>
                <AreaLink href="/" variant="h4" text="AREA" sx={{ ml: 3, textDecoration: "none" }} />
            </AreaBox>
            <AreaBox sx={{ height: "100%", width: "80vw", flexDirection: "row", gap: 4, justifyContent: "end", pr: 2 }}>
                <AreaLink href="/" text="Services" sx={{ fontSize: "1.5em", textDecoration: "none" }} />
                <AreaLink href="/" text="Areas" sx={{ fontSize: "1.5em", textDecoration: "none" }} /> 
                <CreateButton />
                <Avatar sx={{ border: "3px solid black", width: 48, height: 48 }}>P</Avatar>

            </AreaBox>

        </AreaBox>

	);

};

export default Navbar;
