import { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreasTypography, AreaTypography } from "./elements/AreaTypography";
import DeleteIcon from '@mui/icons-material/Delete';

import { useCookies } from "react-cookie";

import Footer from "./Footer";
import Navbar from "./Navbar";
import axios from "axios";
import { AreaButton } from "./elements/AreaButton";
import { IconButton } from "@mui/material";

type Action = {
	[value: string]: string | number;
}

type Reaction = Action;

interface Area {
	id: number;
	area_id: string;
	action: Action;
	reaction: Reaction;
	action_name: string;
	reaction_name: string;
}

export type Parameters = {
    [value: string]: string | number;
}

const Areas: React.FC = () => {

	const [areas, setAreas] = useState<Area[]>([]);

	const [cookies] = useCookies();

	const getAreas = () => {

		const url: string = "http://127.0.0.1:8080/areas";

		const token: string = cookies["token"];
				
		axios.get(url, { 
			headers: {
				"Content-Type": "application/json",
				token
			}
		}).then((res) => {

			console.log(res.data.areas);
			setAreas(res.data.areas);

        });

	};

	const deleteArea = (area_id: string) => {

		const url: string = "http://127.0.0.1:8080/areas";

		const token: string = cookies["token"];
				
		axios.delete(url, { 
			headers: {
				"Content-Type": "application/json",
				token
			},
			data: {
				area_id
			}
		}).then(getAreas);

	}

	useEffect(() => {
		getAreas();
	}, []);

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw" }}>

					<AreaTypography variant="h2" text="Areas" />

					<AreaBox sx={{ height: "100vh", width: "90vw", justifyContent: "start", gap: 3, mt: 2 }}>
						{areas && areas.map((area: Area) => {
							return (
								<AreaBox sx={{ border: "2px solid black", height: "8vh", width: "50vw", flexDirection: "row", borderRadius: 5, backgroundColor: "#d4d4d4", gap: 2 }} key={area.id}>
									<AreaTypography variant="h4" text="If" />
									<AreaTypography text={area.action_name} sx={{ color: "#656565" }} />
									<AreaTypography variant="h4" text="Else" />
									<AreaTypography text={area.reaction_name} sx={{ color: "#656565" }} />
									<IconButton onClick={() => deleteArea(area.area_id)}>
										<DeleteIcon sx={{ color:"#de2626" }} />
									</IconButton>
								</AreaBox>
							);
						})}
					</AreaBox>

				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Areas;
