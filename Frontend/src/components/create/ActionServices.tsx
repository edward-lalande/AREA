import { useEffect, useState } from "react";
import { Action, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaTypography } from "../elements/AreaTypography";
import axios from "axios";
import { Box, Button, Card, Typography } from "@mui/material";

interface ActionServices {
    name: string;
    color: string;
    actions: Action[];
}

type ActionServicesProps = {
    setPage: (value: CreatePage) => void;
    setSelectedActions: (value: ActionServices) => void;
};

const ActionServices: React.FC<ActionServicesProps> = ({
    setPage,
    setSelectedActions
}) => {

    const [actions, setActions] = useState<ActionServices[]>([]);

    const getActions = () => {

        const url = "http://127.0.0.1:8080/actions";

        axios.get(url).then((res) => {

            const actions: ActionServices[] = res.data;

            setActions(actions.filter(element => element !== null));

        });

    }

    useEffect(() => {

		getActions();

	}, []);

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>
            
            <AreaTypography variant="h2" text="Select a service" sx={{ mb: 5 }} />

            <AreaBox sx={{ display: "grid", gridTemplateColumns: "repeat(3, 1fr)", height: "60vh", width: "80vw", marginLeft: 20 }}>

                {actions && actions.map<JSX.Element>((value) => {
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
                                    onClick={() => {setPage(CreatePage.ACTION); setSelectedActions(value)}}
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

export default ActionServices;