import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { Alert, Snackbar } from "@mui/material";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, DiscordButton } from "./elements/AreaButton";

import axios from "axios";
import { useCookies } from "react-cookie";

const LoginForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

	const [open, setOpen] = useState<boolean>(false);
	const [popup, setPopup] = useState<Window | null>(null);

	//eslint-disable-next-line
	const [cookie, setCookie] = useCookies();

	const login = (email: string, password: string) => {

		const url: string = "http://127.0.0.1:8080/user";

		const data = {
			routes: "login",
			mail: email,
			password
		};

		axios.post(url, data).then((res) => {

			setCookie("token", res.data.body.token);
			window.location.href = "/";

		}).catch(() => {
			setOpen(true);
		});
	
	}

	const authDiscord = () => {

		const url: string = "http://127.0.0.1:8083/oauth2";

		axios.get(url).then((res) => {

			const width = 500;
			const height = 600;
			const left = (window.innerWidth - width) / 2;
			const top = (window.innerHeight - height) / 2;
			
			const popup = window.open(
				res.data,
				'_blank',
				`width=${width},height=${height},top=${top},left=${left},resizable=yes,scrollbars=yes,status=yes`
			);

			setPopup(popup);

		});

	}

	useEffect(() => {

		// Trigger l'URL de la popup //

	}, [popup]);

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ width: "30vw", justifyContent: "center", padding: 8 }}>

				<AreaTypography variant="h2" text="AREA" />
				<AreaTypography variant="h4" text="Log in" />

				<AreaBox sx={{ width: "100%", maxWidth: 400, mb: 2 }}>
					<AreaTextField label="Email" onChange={(s) => setEmail(s.target.value)}/>
					<AreaTextField label="Password" type="password" onChange={(s) => setPassword(s.target.value)} />
				</AreaBox>

				<AreaButton text="Log in" onClick={() => login(email, password)}/>

				<AreaTextDivider text="or" />

				<DiscordButton onClick={authDiscord} />

				<AreaBox sx={{ flexDirection: "row", mt: 1 }}>
					<AreaTypography variant="h6" text="New on Area?" sx={{ mr: 2 }} />
					<AreaLink href="/signup" text="Signup here" />
				</AreaBox>

				<Snackbar open={open} autoHideDuration={6000} onClose={() => setOpen(false)}>
					<Alert
						onClose={() => setOpen(false)}
						severity="error"
						variant="filled"
						sx={{ width:"100%" }}
					>
						Login: Invalid email or password
					</Alert>
				</Snackbar>	

			</AreaPaper>

		</AreaBox>

	);

};

export default LoginForm;
