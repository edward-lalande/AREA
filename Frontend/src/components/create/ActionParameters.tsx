import { Button } from "@mui/material";
import { Action, Parameters, CreatePage, Argument } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { AreaButton } from "../elements/AreaButton";
import { AreaTextField } from "../elements/AreaTextFiled";
import { AreasTypography, AreaTypography } from "../elements/AreaTypography";

type ActionParametersProps = {
    action: Action;
    actionParameters: Parameters | undefined;
    setActionParameters: (value: Parameters) => void;
    setPage: (value: CreatePage) => void 
}

const ActionParameters: React.FC<ActionParametersProps> = ({
    action,
    actionParameters,
    setActionParameters,
    setPage
}) => {

    const addActionParameter = (name: string, type: string, value: string) => {

        if (type === "number") {
            setActionParameters(({ ...actionParameters, [name]: Number(value) }));
        }

        if (type === "string") {
            setActionParameters(({ ...actionParameters, [name]: String(value) }));
        }

    };

    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

            <AreaBox sx={{ width: "98vw", flexDirection: "row", justifyContent: "center", gap: 3 }}>
                <Button variant="contained" sx={{ backgroundColor: "black", color: "white", width: "5vw", borderRadius: 5, mb: 3 }} onClick={() => setPage(CreatePage.ACTION)}>Back</Button>
                <AreaTypography variant="h2" text={action.name} sx={{ mb: 5 }} />
            </AreaBox>

            <AreasTypography variant="h4" text={action.description} sx={{ mb: 5 }} />

            { action.arguments && action.arguments.map<JSX.Element>((value: Argument) => {
                
                return (<AreaTextField label={value.display} onChange={(s) => addActionParameter(value.name, value.type, s.target.value) } sx={{ width: "20vw" }} />);

            })}

            <AreaButton text="Validate" onClick={() => { setPage(CreatePage.CREATE); }}/>

        </AreaBox>

    );

}

export default ActionParameters;