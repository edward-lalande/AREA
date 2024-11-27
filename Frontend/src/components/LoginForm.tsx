import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, GoogleButton } from "./elements/AreaButton";
import { useState } from "react";
import axios from "axios";

async function loginMe(email: string, password: string) {
	const url = "http://127.0.0.1:8080/user"
	const body = {
		routes: "login",
		mail: email,
		password
	}
	axios.post(url, body).then((rep) => {
		console.log(rep)
		window.location.href = "/";
	}).catch((e) => {
		console.error(e)
	})

}


const LoginForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ width: "30vw", justifyContent: "center", padding: 8 }}>

				<AreaTypography variant="h2" text="AREA" />
				<AreaTypography variant="h4" text="Log in" />

				<AreaBox sx={{ width: "100%", maxWidth: 400, mb: 2 }}>
					<AreaTextField label="Email" onChange={(s) => setEmail(s.target.value)}/>
					<AreaTextField label="Password" type="password" onChange={(s) => setPassword(s.target.value)} />
				</AreaBox>

				<AreaButton text="Log in" onClick={() => loginMe(email, password)}/>

				<AreaTextDivider text="or" />

				<GoogleButton />

				<AreaBox sx={{ flexDirection: "row", mt: 1 }}>
					<AreaTypography variant="h6" text="New on Area?" sx={{ mr: 2 }} />
					<AreaLink href="/signup" text="Signup here" />
				</AreaBox>

			</AreaPaper>

		</AreaBox>

	);

};

export default LoginForm;
