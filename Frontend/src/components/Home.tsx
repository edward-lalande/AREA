import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";

const Home: React.FC = () => {

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

			</AreaPaper>

		</AreaBox>

	);

};

export default Home;
