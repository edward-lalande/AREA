import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";

import Navbar from "./Navbar";

const Home: React.FC = () => {

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw", padding: 1 }}>

				<Navbar />

			</AreaPaper>

		</AreaBox>

	);

};

export default Home;
