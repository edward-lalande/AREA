import { Action, CreatePage } from "../Create";
import { AreaBox } from "../elements/AreaBox";
import { ServiceButton } from "../elements/AreaButton";
import { AreaTypography } from "../elements/AreaTypography";

type ActionElementProps = {
    selectedActions: Action[];
    setPage: (value: CreatePage) => void;
    setAction: (value: Action) => void;
}

const ActionElement: React.FC<ActionElementProps> = ({
    selectedActions,
    setPage,
    setAction
}) => {
    
    return (

        <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

            <AreaTypography variant="h2" text="Select an action" sx={{ mb: 5 }} />

            <AreaBox sx={{ height: "20vh", width: "98vw", gap: 2, flexDirection: "row" }}>

                {selectedActions && selectedActions.map<JSX.Element>((value: Action) => {
                    return (<ServiceButton text={value.name} backgroundColor="#000" onClick={() => {setPage(CreatePage.ACTION_ARGUMENTS); setAction(value)}}/>);
                })}

            </AreaBox>

        </AreaBox>

    );

}

export default ActionElement;