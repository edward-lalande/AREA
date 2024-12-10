import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";

import Footer from "./Footer";
import Navbar from "./Navbar";
import axios from "axios";

interface Service {
	name: string;
}

const Home: React.FC = () => {

	const [services, setServices] = useState<Service[]>([]);

	const getServices = () => {

		const url: string = "http://127.0.0.1:8080/services";

		axios.get(url).then((res) => {

            setServices(res.data);

        });

	};

	useEffect(() => {

		getServices();

	}, [services]);

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw" }}>

					<AreaTypography variant="h2" text="Services" />

					{services && services.map<JSX.Element>((service: Service) => {
						return (<p>service.name</p>);
					})}

				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Home;
