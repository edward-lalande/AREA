import { Avatar, Typography } from "@mui/material";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";

import Navbar from "./Navbar";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaButton } from "./elements/AreaButton";

const Account: React.FC = () => {

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw", padding: 1 }}>

				<Navbar />

                <AreaBox sx={{ heigth: "80vh", width: "30vw", mt: 5, gap: 1.5 }}>
                    
                    <AreaTypography variant="h3" text="Account settings" />

                    <AreaDivider />

                    <AreaTypography variant="h4" text="Profile" sx={{ mt: 2 }} />

                    <Avatar sx={{ border: "4px solid black", width: 80, height: 80 }}>P</Avatar>

                    <AreaTypography variant="h4" text="Account" sx={{ mt: 2 }} />

                    <AreaTextField label="Name" />
                    <AreaTextField label="Lastname" />
                    <AreaTextField label="Email" />

                    <AreaButton text="Update" />

                </AreaBox>

			</AreaPaper>

		</AreaBox>

	);

};

export default Account;
