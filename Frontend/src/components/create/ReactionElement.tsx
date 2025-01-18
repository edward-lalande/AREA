import { Button, Card, Typography } from "@mui/material";
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

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

            <AreaTypography variant="h2" text="Select a reaction" sx={{ mb: 5 }} />

            <AreaBox sx={{ height: "20vh", width: "98vw", gap: 2, flexDirection: "row" }}>

                {selectedReactions && selectedReactions.reactions.map<JSX.Element>((value: Reaction) => {
                    return (
                        <Card variant="outlined" sx={{ width: "40vw", height: "6vh", borderRadius: 5 }}>
                            <AreaBox sx={{ width: "40vw", height: "6vh", gap: 2, flexDirection: "row", textAlign: "center", justifyContent: "center" }}>
                                <Typography color={selectedReactions?.color} sx={{ fontSize: 24, mb: 0.5 }}>
                                    {value.name}
                                </Typography>
                                <Button
                                    variant="contained"
                                    onClick={() => {setPage(CreatePage.REACTION_ARGUMENTS); setReaction(value)}}
                                    sx={{
                                        backgroundColor: selectedReactions?.color,
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

        </AreaBox>

    );

}

export default ReactionElement;