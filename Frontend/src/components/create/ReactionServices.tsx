import { useEffect, useState } from "react";
import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaTypography } from "../elements/AreaTypography";
import axios from "axios";
import { Box, Button, Card, Typography } from "@mui/material";

interface ReactionServices {
    name: string;
    color: string;
    reactions: Reaction[];
}

type ReactionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedReactions: (value: ReactionServices) => void;
}

const ReactionServices: React.FC<ReactionServicesProps> = ({
    setPage,
    setSelectedReactions
}) => {

    const [reactions, setReactions] = useState<ReactionServices[]>([]);

    const getReaction = () => {

        const url = "http://127.0.0.1:8080/reactions";

        axios.get(url).then((res) => {

            const reactions: ReactionServices[] = res.data;

            setReactions(reactions.filter(element => element !== null));

        });

    }

    useEffect(() => {

		getReaction();

	}, []);

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

           <AreaBox sx={{ width: "98vw", flexDirection: "row", justifyContent: "center", gap: 3 }}>
                <Button variant="contained" sx={{ backgroundColor: "black", color: "white", width: "5vw", borderRadius: 5, mb: 4 }} onClick={() => setPage(CreatePage.CREATE)}>Back</Button>
                <AreaTypography variant="h2" text="Select a service" sx={{ mb: 5 }} />
            </AreaBox>

            <AreaBox sx={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", height: "60vh", width: "80vw", marginLeft: 20 }}>

                {reactions && reactions.map<JSX.Element>((value) => {
                    return (
                        <Card variant="outlined" sx={{ width: "20vw", height: "16vh", display: "flex", flexDirection: "row", borderRadius: 5 }}>
                            <AreaBox sx={{ width: "8vw", height: "16vh" }}>
                                <img src={value.name.toLowerCase().replace(' ', '_') + ".png"} width={80} height={80} />
                            </AreaBox>
                            <Box sx={{ width: "12vw", height: "16vh", backgroundColor: value.color, padding: 2, gap: 2 }}>
                                <Typography color="white" sx={{ fontSize: 24, mb: 0.5 }}>
                                    {value.name}
                                </Typography>
                                <Button
                                    variant="contained"
                                    onClick={() => {setPage(CreatePage.REACTION); setSelectedReactions(value)}}
                                    sx={{
                                        backgroundColor: "white",
                                        color: "black",
                                        textTransform: "none",
                                        borderRadius: 3,
                                        mt: 1,
                                    }}
                                >
                                    Choose
                                </Button>
                            </Box>
                        </Card>
                    );
                })}

            </AreaBox>

        </AreaBox>

    );

}

export default ReactionServices;