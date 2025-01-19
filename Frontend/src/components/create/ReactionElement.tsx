import { Box, Button, Card, Typography } from "@mui/material";
import { Reaction, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";
import ReactionServices from "./ReactionServices";


type ReactionElementProps = {
    selectedReactions: ReactionServices | null;
    setPage: (value: CreatePage) => void;
    setReaction: (value: Reaction) => void;
}

const ReactionElement: React.FC<ReactionElementProps> = ({
    selectedReactions,
    setPage,
    setReaction
}) => {

    return (

        <Box sx={{ display: "flex", flexDirection: "column", alignItems: "center", width: "70vw", backgroundColor: selectedReactions?.color, borderRadius: 5, marginTop: "5vh", marginBottom: "5vh" }}>

            <AreaBox sx={{ width: "70vw", flexDirection: "row", justifyContent: "center", gap: 3 }}>
                <Button variant="contained" sx={{ backgroundColor: "white", color: selectedReactions?.color, width: "5vw", borderRadius: 5, mt: 0.5 }} onClick={() => setPage(CreatePage.REACTION_SERVICES)}>Back</Button>
                <AreaTypography variant="h2" text="Select a reaction" sx={{ marginTop: 2, color: "white" }} />
            </AreaBox>

            <AreaBox sx={{ height: "60vh", width: "60vw", gap: 2 }}>

                {selectedReactions && selectedReactions.reactions.map<JSX.Element>((value: Reaction) => {
                    return (
                        <Card variant="outlined" sx={{ width: "40vw", height: selectedReactions.reactions.length > 7 ? "4vh" : "6vh", borderRadius: 5 }}>
                            <AreaBox sx={{ width: "40vw", height: selectedReactions.reactions.length > 7 ? "4vh" : "6vh", gap: 2, flexDirection: "row", textAlign: "center", justifyContent: "center" }}>
                                <Typography color={selectedReactions?.color} sx={{ fontSize: selectedReactions.reactions.length > 7 ? 18 : 24, mb: 0.5 }}>
                                    {value.name}
                                </Typography>
                                <Button
                                    variant="contained"
                                    onClick={() => {setPage(CreatePage.REACTION_ARGUMENTS); setReaction(value)}}
                                    sx={{
                                        backgroundColor: selectedReactions?.color,
                                        color: "white",
                                        textTransform: "none",
                                        borderRadius: 3,
                                        height: "65%"
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

export default ReactionElement;