import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { Alert, Snackbar } from "@mui/material";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, DiscordButton, GithubButton, GitlabButton, GoogleButton, SpotifyButton } from "./elements/AreaButton";

import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { useSearchParams } from "react-router-dom";

import axios from "axios";

import { OauthServices } from "./LoginForm";
import { Code } from "./LoginForm";

const SignupForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [name, setName] = useState<string>("");
    const [lastname, setLastname] = useState<string>("");
    const [password, setPassword] = useState<string>("");

	const [open, setOpen] = useState<boolean>(false);

	//eslint-disable-next-line
	const [cookie, setCookie] = useCookies();

	const [searchParams] = useSearchParams();

	const signup = (email: string, name: string, lastname: string, password: string) => {

		const url: string = "http://127.0.0.1:8080/sign-up";

		const data = {
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

	const oauth = (service: OauthServices) => {
		const url: string = `http://127.0.0.1:8080/${service}/oauth`;

		axios.get(url).then((res) => {

			window.location.href = res.data;

		});

	}

	const getToken = (code: string, service: OauthServices) => {

		const url: string = `http://127.0.0.1:8080/${service}/access-token`;

		axios.post(url, { code }).then((res) => {

			if (res.data.body != undefined) {
				setCookie("token", res.data.body);
			}
			window.location.href = "/";

		});

		
	}

	useEffect(() => {

		const codes: Code[] = [];

		codes.push({ name: "discord_code", service: OauthServices.DISCORD });
		codes.push({ name: "gitlab_code", service: OauthServices.SPOTIFY });
		codes.push({ name: "github_code", service: OauthServices.GITHUB });
		codes.push({ name: "google_code", service: OauthServices.GOOGLE });

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
				<AreaTypography variant="h4" text="Sign up" />

				<AreaBox sx={{ width: "100%", maxWidth: 400, mb: 2 }}>
					<AreaTextField label="Email" onChange={(s) => setEmail(s.target.value)} />
					<AreaTextField label="Password" type="password" onChange={(s) => setPassword(s.target.value)}/>
					<AreaTextField label="Name" onChange={(s) => setName(s.target.value)}/>
					<AreaTextField label="Lastname" onChange={(s) => setLastname(s.target.value)}/>
				</AreaBox>

				<AreaButton text="Sign up" onClick={() => signup(email, name, lastname, password)}/>

				<AreaTextDivider text="or" />

				<DiscordButton text="Continue with Discord" onClick={() => oauth(OauthServices.DISCORD)} />
				<GoogleButton text="Continue with Google" onClick={() => oauth(OauthServices.GOOGLE)} />
				<GithubButton text="Continue with Github" onClick={() => oauth(OauthServices.GITHUB)} />
				<GitlabButton text="Continue with Gitlab" onClick={() => oauth(OauthServices.GITLAB)} />

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
