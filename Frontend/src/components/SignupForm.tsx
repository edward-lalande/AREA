import { AreaBox } from "./elements/AreaBox";
import { AreaLink } from "./elements/AreaLink";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTextDivider } from "./elements/AreaDivider";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaTypography } from "./elements/AreaTypography";
import { AreaButton, GoogleButton } from "./elements/AreaButton";

const SignupForm: React.FC = () => {

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ width: "30vw", justifyContent: "center", padding: 8 }}>

				<AreaTypography variant="h2" text="AREA" />
				<AreaTypography variant="h4" text="Sign up" />

				<AreaBox sx={{ width: "100%", maxWidth: 400, mb: 2 }}>
					<AreaTextField label="Email" />
					<AreaTextField label="Password" type="password" />
					<AreaTextField label="Name" />
					<AreaTextField label="Surname" />
				</AreaBox>

				<AreaButton text="Sign up" />

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
