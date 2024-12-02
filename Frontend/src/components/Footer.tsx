import { AreaBox } from "./elements/AreaBox";
import { Logout } from "./elements/AreaButton";

import { useCookies } from "react-cookie";

const Footer: React.FC = () => {

    //eslint-disable-next-line
    const [cookie, setCookie, removeCookie] = useCookies();

    const logout = () => {
        removeCookie("token");
        window.location.href = "/login";
    }

	return (

        <AreaBox sx={{ height: "8vh", width: "100%", flexDirection: "row", mb: 1 }}>

            <AreaBox sx={{ height: "100%", width: "100vw", flexDirection: "row", gap: 4, justifyContent: "end", pr: 3 }}>

                <Logout onClick={logout} />

            </AreaBox>

        </AreaBox>

	);

};

export default Footer;