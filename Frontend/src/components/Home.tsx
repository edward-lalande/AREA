import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";

import Footer from "./Footer";
import Navbar from "./Navbar";
import axios from "axios";
import { ServiceButton } from "./elements/AreaButton";

interface Service {
	id: number;
	name: string;
	color: string;
}

const Home: React.FC = () => {

	const [services, setServices] = useState<Service[]>([]);

	const getServices = () => {

		const url: string = "http://127.0.0.1:8080/services";

		axios.get(url).then((res) => {

			console.log(res.data);
			setServices(res.data.services);

        });

	};

	useEffect(() => {
		getServices();
	}, []);

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw" }}>

					<AreaTypography variant="h2" text="Services" />

					<AreaBox sx={{ display: "grid", gridTemplateColumns: "repeat(5, 1fr)", height: "50vh", width: "90vw", gap: 3, ml: 18, mt: 5 }}>

						{services && services.map<JSX.Element>((service: Service) => {
							return (<ServiceButton key={service.id} text={service.name} backgroundColor={service.color} />);
						})}

					</AreaBox>

				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Home;
