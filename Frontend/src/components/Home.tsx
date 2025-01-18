import React, { useEffect, useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";

import Footer from "./Footer";
import Navbar from "./Navbar";
import axios from "axios";
import { ServiceCard } from "./elements/AreaButton";
import { Box, Button, Card, Dialog, Typography } from "@mui/material";

interface Service {
	id: number;
	name: string;
	color: string;
	description: string;
}

interface Argument {
    display: string;
    name: string;
    type: string;
}

interface Action {
    name: string;
    description: string;
    action_id: number;
    action_type: number;
    arguments: Argument[];
}

interface Reaction {
    name: string;
    description: string;
    reaction_id: number;
    reaction_type: number;
    arguments: Argument[];
}

interface ActionServices {
	name: string;
    actions: Action[];
}

interface ReactionServices {
	name: string;
	reactions: Reaction[];
}

const Home: React.FC = () => {
	
	const [selectedService, setSelectedService] = useState<Service | null>(null);
	const [selectedActions, setSelectedActions] = useState<string[]>([]);
	const [selectedReactions, setSelectedReactions] = useState<string[]>([]);

	const setActions = (actionServices: ActionServices[]) => {

		let actions: string[] = [];

		for (let i = 0; actionServices && i < actionServices.length; i++) {
			if (actionServices[i]?.name === selectedService?.name) {
				for (let j = 0; j < actionServices[i].actions.length; j++) {
					actions.push(actionServices[i].actions[j].name);
				}
			}
		}

		setSelectedActions(actions);
	}

	const setReactions = (reactionServices: ReactionServices[]) => {

		let reactions: string[] = [];

		for (let i = 0; reactionServices && i < reactionServices.length; i++) {
			if (reactionServices[i]?.name === selectedService?.name) {
				for (let j = 0; j < reactionServices[i].reactions.length; j++) {
					reactions.push(reactionServices[i].reactions[j].name);
				}
			}
		}

		setSelectedReactions(reactions);
	}

	const [services, setServices] = useState<Service[]>([]);

	const getServices = () => {

		const url: string = "http://127.0.0.1:8080/services";

		axios.get(url).then((res) => {

			setServices(res.data.services);

        });

	};

	const getActions = () => {

		const url: string = "http://127.0.0.1:8080/actions";

		axios.get(url).then((res) => {
			let actions: ActionServices[] = res.data;

			actions = actions.filter(element => element !== null);
			setActions(actions);
		});

	}

	const getReactions = () => {

		const url: string = "http://127.0.0.1:8080/reactions";

		axios.get(url).then((res) => {
			let reactions: ReactionServices[] = res.data;
			reactions = reactions.filter(element => element !== null);
			setReactions(reactions);
		});

	}

	useEffect(() => {
		getServices();
		getActions();
		getReactions();
	}, [selectedService]);

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw" }}>

					<AreaTypography variant="h2" text="Services" />

					<AreaBox sx={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", height: "60vh", width: "94vw", gap: 1, mt: 3 }}>

						{services && services.map<JSX.Element>((service: Service) => {
							return (
								<ServiceCard service={service} onClick={setSelectedService} />
							);
						})}

						<Dialog open={selectedService != null} sx={{ backdropFilter: "blur(3px)" }} PaperProps={{ style: { borderRadius: 25 }}}>
							<AreaBox sx={{ width: "30vw", height: "100vh", display: "flex", backgroundColor: selectedService?.color }}>
								<Box sx={{ width: "28vw", height: "35vh", padding: 2, textAlign: "center", backgroundColor: "white", m: 3, borderRadius: 5 }}>
									<Typography sx={{ fontSize: 24, mb: 0.5, color: selectedService?.color }}>
										{selectedService?.name}
									</Typography>
									<img src={selectedService?.name.toLowerCase().replace(' ', '_') + ".png"} width={80} height={80} />
									<Typography sx={{ fontSize: 16, color: selectedService?.color }}>
										{selectedService?.description}
									</Typography>
								</Box>
								<Box sx={{ width: "30vw", height: "65vh", backgroundColor: selectedService?.color, paddingLeft: 3, gap: 3 }}>
									{selectedActions.length > 0 && <Typography sx={{ fontSize: 24, mb: 0.5, color: "white" }}>Action</Typography>}
									{selectedActions && selectedActions.map(value => (
										<Card variant="outlined" sx={{ width: "28vw", height: "4vh", borderRadius: 2, mt: 1, justifyContent: "center", padding: 0.2, paddingLeft: 1 }}>
											<Typography color={selectedService?.color} sx={{ fontSize: 18, mb: 0.5 }}>{value}</Typography>
										</Card>
									))}
									{selectedReactions.length > 0 && <Typography sx={{ fontSize: 24, mb: 0.5, color: "white", mt: 2 }}>Reactions</Typography>}
									{selectedReactions && selectedReactions.map(value => (
										<Card variant="outlined" sx={{ width: "28vw", height: "4vh", borderRadius: 2, mt: 1, justifyContent: "center", padding: 0.2, paddingLeft: 1 }}>
											<Typography color={selectedService?.color} sx={{ fontSize: 18, mb: 0.5 }}>{value}</Typography>
										</Card>
									))}
									<Button
										variant="contained"
										onClick={() => setSelectedService(null)}
										sx={{
											backgroundColor: "white",
											color: selectedService?.color,
											textTransform: "none",
											borderRadius: 3,
											mt: selectedReactions.length ? 3 : selectedActions.length ? 3 : 0
										}}
									>
										Close
									</Button>
								</Box>
							</AreaBox>
						</Dialog>

					</AreaBox>

				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Home;
