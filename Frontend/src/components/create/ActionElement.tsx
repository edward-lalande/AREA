import { Box, Button, Card, Typography } from "@mui/material";
import { Action, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaTypography } from "../elements/AreaTypography";
import ActionServices from "./ActionServices";

type ActionElementProps = {
    selectedActions: ActionServices | null;
    setPage: (value: CreatePage) => void;
    setAction: (value: Action) => void;
}

const ActionElement: React.FC<ActionElementProps> = ({
    selectedActions,
    setPage,
    setAction
}) => {
    
    return (

        <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center", width: "70vw", backgroundColor: selectedActions?.color, borderRadius: 5, marginTop: "5vh", marginBottom: "5vh" }}>

            <AreaTypography variant="h2" text="Select an action" sx={{ marginTop: 2, color: "white" }} />

            <AreaBox sx={{ height: "60vh", width: "60vw", gap: 2 }}>

                {selectedActions && selectedActions.actions.map<JSX.Element>((value: Action) => {
                    return (
                        <Card variant="outlined" sx={{ width: "40vw", height: "6vh", borderRadius: 5 }}>
                            <AreaBox sx={{ width: "40vw", height: "6vh", gap: 2, flexDirection: "row", textAlign: "center", justifyContent: "center" }}>
                                <Typography color={selectedActions?.color} sx={{ fontSize: 24, mb: 0.5 }}>
                                    {value.name}
                                </Typography>
                                <Button
                                    variant="contained"
                                    onClick={() => {setPage(CreatePage.ACTION_ARGUMENTS); setAction(value)}}
                                    sx={{
                                        backgroundColor: selectedActions?.color,
                                        color: "white",
                                        textTransform: "none",
                                        borderRadius: 3
                                    }}
                                >
                                    Choose
                                </Button>
                            </AreaBox>
                        </Card>
                    );
                })}

            </AreaBox>

        </Box>

    );

}

export default ActionElement;