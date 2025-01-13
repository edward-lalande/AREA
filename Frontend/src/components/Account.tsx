import { AreaBox } from "./elements/AreaBox";
import { AccountButton, AreaButton, AsanaButton, DiscordButton, DropboxButton, GithubButton, GitlabButton, GoogleButton, MiroButton, SpotifyButton } from "./elements/AreaButton";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreasTypography, AreaTypography } from "./elements/AreaTypography";
import DoneOutlineIcon from '@mui/icons-material/DoneOutline';
import CancelIcon from '@mui/icons-material/Cancel';
import Footer from "./Footer";

import { useCookies } from "react-cookie";
import { useSearchParams } from "react-router-dom";

import Navbar from "./Navbar";
import { useEffect, useState } from "react";
import axios from "axios";
import { OauthServices } from "./LoginForm";
import { Code } from "./LoginForm";

type User = {
	id: string;
	mail: string;
	password: string;
	name: string;
	lastname: string;
	asana_token: string;
	discord_token: string;
	dropbox_token: string;
	github_token: string;
	gitlab_token: string;
	google_token: string;
	miro_token: string;
	spotify_token: string;
}

const Account: React.FC = () => {

	const [user, setUser] = useState<User>();
	const [cookies] = useCookies();
	const [email, setEmail] = useState<string>("");
	const [name, setName] = useState<string>("");
	const [lastname, setLastname] = useState<string>("");

	const [searchParams] = useSearchParams();

	const oauth = (service: OauthServices) => {

		console.log(service)

		//const url: string = `http://127.0.0.1:8080/${service}/add-oauth`;
//
		//axios.get(url, {
		//	headers: {
		//		"Content-Type": "application/json",
		//		"token": cookies["token"]
		//	}
		//}).then((res) => {
//
		//	window.location.href = res.data;
//
		//});

	}

	const getToken = (code: string, service: OauthServices) => {

		const url: string = `http://127.0.0.1:8080/${service}/add-access-token`;

		axios.post(url, { code }, {
			headers: {
				"Content-Type": "application/json",
				"token": cookies["token"]
			}
		}).then(getUser);

	}

	const getUser = () => {

		const url: string = "http://127.0.0.1:8085/user";

		const token: string = cookies["token"];

		axios.get(url, {
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then((res) => {
			setUser(res.data.user);
			setEmail(res.data.user.mail);
			setName(res.data.user.name)
			setLastname(res.data.user.lastname)
		})

	}

	const editEmail = () => {

		const url: string = "http://127.0.0.1:8085/update-email";

		const token: string = cookies["token"];

		const data = {
			"mail": email
		}

		axios.post(url, data, {
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then(getUser)

	}

	const editName = () => {

		const url: string = "http://127.0.0.1:8085/update-name";

		const token: string = cookies["token"];

		const data = {
			"name": name
		}

		axios.post(url, data, {
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then(getUser)

	}

	const editLastname = () => {

		const url: string = "http://127.0.0.1:8085/update-lastname";

		const token: string = cookies["token"];

		const data = {
			"lastname": lastname
		}

		axios.post(url, data, {
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then(getUser)

	}

	useEffect(() => {

		const codes: Code[] = [];

		codes.push({ name: "discord_code", service: OauthServices.DISCORD });
		codes.push({ name: "spotify_code", service: OauthServices.SPOTIFY });
		codes.push({ name: "github_code", service: OauthServices.GITHUB });
		codes.push({ name: "gitlab_code", service: OauthServices.GITLAB });
		codes.push({ name: "google_code", service: OauthServices.GOOGLE });
		codes.push({ name: "dropbox_code", service: OauthServices.DROPBOX });
		codes.push({ name: "asana_code", service: OauthServices.ASANA });
		codes.push({ name: "miro_code", service: OauthServices.MIRO });

		for (let i = 0; i < codes.length; i++) {

			const findCode: string | null = searchParams.get(codes[i].name);

			if (findCode) {
				getToken(findCode, codes[i].service);
			}

		}

		getUser();

	}, [searchParams]);

	return (

        <AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "80vw", flexDirection: "row", justifyContent: "space-around" }}>

					<AreaBox sx={{ height: "80vh", width: "20vw" }}>
						<AreaTypography variant="h2" text="Account" sx={{ mb: 10 }} />

						<AreasTypography text={"Email: " + user?.mail} />
						<AreaBox sx={{ height: "10vh", width: "20vw", flexDirection: "row", gap: 5 }}>
							<AreaTextField onChange={(s) => setEmail(s.target.value)}/>
							<AccountButton text="Edit" onClick={editEmail}/>
						</AreaBox>

						<AreasTypography text={"Name: " + user?.name} />
						<AreaBox sx={{ height: "10vh", width: "20vw", flexDirection: "row", gap: 5 }}>
							<AreaTextField onChange={(s) => setName(s.target.value)} />
							<AccountButton text="Edit" onClick={editName}/>
						</AreaBox>

						<AreasTypography text={"Lastname: " + user?.lastname} />
						<AreaBox sx={{ height: "10vh", width: "20vw", flexDirection: "row", gap: 5 }}>
							<AreaTextField onChange={(s) => setLastname(s.target.value)} />
							<AccountButton text="Edit" onClick={editLastname}/>
						</AreaBox>
					</AreaBox>

					<AreaBox sx={{ height: "80vh", width: "40vw", gap: 2 }}>
						<AreaTypography variant="h3" text="Services" />
						<AreaBox sx={{ width: "40vw", flexDirection: "row", gap: 2 }}>
							<AreaBox sx={{ gap: 2 }}>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<DiscordButton text="Connect Discord" onClick={() => oauth(OauthServices.DISCORD)} />
									{user?.discord_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<SpotifyButton text="Connect Spotify" onClick={() => oauth(OauthServices.SPOTIFY)} />
									{user?.spotify_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<GithubButton text="Connect Github" onClick={() => oauth(OauthServices.GITHUB)} />
									{user?.github_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<GitlabButton text="Connect Gitlab" onClick={() => oauth(OauthServices.GITLAB)} />
									{user?.gitlab_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
							</AreaBox>
							<AreaBox sx={{ gap: 2 }}>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<GoogleButton text="Connect Google" onClick={() => oauth(OauthServices.GOOGLE)} />
									{user?.google_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<DropboxButton text="Connect Dropbox" onClick={() => oauth(OauthServices.DROPBOX)} />
									{user?.dropbox_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<AsanaButton text="Connect Asana" onClick={() => oauth(OauthServices.ASANA)} />
									{user?.asana_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
								<AreaBox sx={{ flexDirection: "row", gap: 2 }}>
									<MiroButton text="Connect Miro" onClick={() => oauth(OauthServices.MIRO)} />
									{user?.miro_token ? <DoneOutlineIcon sx={{ color: "green" }} /> : 
									<CancelIcon sx={{ color: "red" }} />}
								</AreaBox>
							</AreaBox>
						</AreaBox>
					</AreaBox>


				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Account;