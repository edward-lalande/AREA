import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import Footer from "./Footer";

import Navbar from "./Navbar";

const Account: React.FC = () => {

	return (

        <AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

                <Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw" }} />

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Account;