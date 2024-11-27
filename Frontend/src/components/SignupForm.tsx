import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, GoogleButton } from "./elements/AreaButton";
import { useState } from "react";
import axios from "axios";

function signUpMe(email: string, name: string, lastname: string, password: string) {
	const url = "http://127.0.0.1:8080/user"
	const body = {
		routes: "sign-up",
		mail: email,
		name,
		lastname,
		password
	}
	axios.post(url, body).then((rep) => {
		console.log(rep)
		window.location.href = "/"
	}).catch((e) => {
		console.error(e)
	})
}

const SignupForm: React.FC = () => {

    const [email, setEmail] = useState<string>("");
    const [name, setName] = useState<string>("");
    const [lastname, setLastname] = useState<string>("");
    const [password, setPassword] = useState<string>("");

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

				<AreaButton text="Sign up" onClick={() => signUpMe(email, name, lastname, password)}/>

				<AreaTextDivider text="or" />

				<GoogleButton />

				<AreaBox sx={{ flexDirection: "row", mt: 1 }}>
					<AreaTypography variant="h6" text="Already on AREA?" sx={{ mr: 2 }} />
					<AreaLink href="/login" text="Login here" />
				</AreaBox>

			</AreaPaper>

		</AreaBox>

	);

};

export default SignupForm;
