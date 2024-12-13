import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { Alert, Snackbar } from "@mui/material";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, DiscordButton } from "./elements/AreaButton";

import { useState } from "react";
import { useCookies } from "react-cookie";

import axios from "axios";

const SignupForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [name, setName] = useState<string>("");
    const [lastname, setLastname] = useState<string>("");
    const [password, setPassword] = useState<string>("");

	const [open, setOpen] = useState<boolean>(false);

	//eslint-disable-next-line
	const [cookie, setCookie] = useCookies();

	const signup = (email: string, name: string, lastname: string, password: string) => {

		const url: string = "http://127.0.0.1:8080/user";

		const data = {
			routes: "sign-up",
			mail: email,
			password,
			name,
			lastname
		};

		axios.post(url, data).then((res) => {

			setCookie("token", res.data.body.token);
			window.location.href = "/login";

		}).catch(() => {
			setOpen(true);
		});
	
	}

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ width: "30vw", justifyContent: "center", padding: 8 }}>

				<AreaTypography variant="h2" text="AREA" />
				<AreaTypography variant="h4" text="Sign up" />

				<AreaBox sx={{ width: "100%", maxWidth: 400, mb: 2 }}>
					<AreaTextField label="Email" onChange={(s) => setEmail(s.target.value)} />
					<AreaTextField label="Password" type="password" onChange={(s) => setPassword(s.target.value)}/>
					<AreaTextField label="Name" onChange={(s) => setName(s.target.value)}/>
					<AreaTextField label="Lastname" onChange={(s) => setLastname(s.target.value)}/>
				</AreaBox>

				<AreaButton text="Sign up" onClick={() => signup(email, name, lastname, password)}/>

				<AreaTextDivider text="or" />

				<DiscordButton />

				<AreaBox sx={{ flexDirection: "row", mt: 1 }}>
					<AreaTypography variant="h6" text="Already on AREA?" sx={{ mr: 2 }} />
					<AreaLink href="/login" text="Login here" />
				</AreaBox>

				<Snackbar open={open} autoHideDuration={6000} onClose={() => setOpen(false)}>
					<Alert
						onClose={() => setOpen(false)}
						severity="error"
						variant="filled"
						sx={{ width:"100%" }}
					>
						Signup: User already exist
					</Alert>
				</Snackbar>	

			</AreaPaper>

		</AreaBox>

	);

};

export default SignupForm;
