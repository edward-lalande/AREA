import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { Alert, Snackbar } from "@mui/material";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, DiscordButton, GithubButton, GitlabButton, SpotifyButton, GoogleButton, DropboxButton } from "./elements/AreaButton";

import axios from "axios";
import { useCookies } from "react-cookie";
import { useSearchParams } from "react-router-dom";

enum OauthServices {
	DISCORD = "discord",
	SPOTIFY = "spotify",
	GITHUB = "github",
	GITLAB = "gitlab",
	GOOGLE = "google",
	DROPBOX = "dropbox"
}

type Code = {
	name: string;
	service: OauthServices;
}

const LoginForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

	const [open, setOpen] = useState<boolean>(false);

	//eslint-disable-next-line
	const [_, setCookie] = useCookies();

	const [searchParams] = useSearchParams();

	const login = (email: string, password: string) => {

		const url: string = "http://127.0.0.1:8080/login";

		const data = {
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

	const oauth = (service: OauthServices) => {

		const url: string = `http://127.0.0.1:8080/${service}/oauth`;

		axios.get(url).then((res) => {

			window.location.href = res.data;

		});

	}

	const getToken = (code: string, service: OauthServices) => {

		const url: string = `http://127.0.0.1:8080/${service}/access-token`;

		axios.post(url, { code }).then((res) => {
			setCookie("token", res.data.body.access_token);
			window.location.href = "/";
		});
		
	}

	useEffect(() => {

		const codes: Code[] = [];

		codes.push({ name: "discord_code", service: OauthServices.DISCORD });
		codes.push({ name: "spotify_code", service: OauthServices.SPOTIFY });
		codes.push({ name: "github_code", service: OauthServices.GITHUB });
		codes.push({ name: "gitlab_code", service: OauthServices.GITLAB });
		codes.push({ name: "google_code", service: OauthServices.GOOGLE });
		codes.push({ name: "dropbox_code", service: OauthServices.DROPBOX });

		for (let i = 0; i < codes.length; i++) {

			const findCode: string | null = searchParams.get(codes[i].name);

			if (findCode) {
				getToken(findCode, codes[i].service);
			}

		}

	}, [searchParams]);

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

				<DiscordButton onClick={() => oauth(OauthServices.DISCORD)} />
				<SpotifyButton onClick={() => oauth(OauthServices.SPOTIFY)} />
				<GithubButton onClick={() => oauth(OauthServices.GITHUB)} />
				<GitlabButton onClick={() => oauth(OauthServices.GITLAB)} />
				<GoogleButton onClick={() => oauth(OauthServices.GOOGLE)} />
				<DropboxButton onClick={() => oauth(OauthServices.DROPBOX)} />

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
