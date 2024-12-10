import { useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AddButton, ServiceButton } from "./elements/AreaButton";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";
import Footer from "./Footer";

import Navbar from "./Navbar";
import { AreaButton } from "./elements/AreaButton";

import axios from "axios";
import { AreaTextField } from "./elements/AreaTextFiled";

const services: string[] = [
    "Time"
];

const actions: string[] = [
    "Every day at"
];

const reactions: string[] = [
    "Send a message on channel"
];

const Create: React.FC = () => {

    const [action, setAction] = useState<string>("");
    const [reaction, setReaction] = useState<string>("");

    const [page, setPage] = useState("Create");

    const [hour, setHour] = useState<number>(0);
    const [minute, setMinute] = useState<number>(0);

    const [channel, setChannel] = useState<string>("");
    const [message, setMessage] = useState<string>("");

    const createArea = (hour: number, minute: number, channel: string, message: string) => {

        const url: string = "http://127.0.0.1:8080/area";

		const data = [{
            user_token: "AREA",
            action: {
                action_id: 2,
                action_type: 0,
                continent: "Europe",
                city: "Paris",
                hour,   
                minute
            },
            reaction: {
                reaction_id: 3,
                reaction_type: 0,
                channel_id: channel,
                message
            }
		}];

        axios.post(url, data).then(() => {

            window.location.href = "/";

        });

    };

	return (

        <AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

                <Navbar />

                { page === "Create" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Create an area" sx={{ mb: 5 }} />

                    <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "#000", width: "50vw", padding: "2%", borderRadius: 5 }}>
                        <AreaTypography variant="h3" text={action ? "If" : "If this"} color="white" />
                        {action ? <AreaTypography text={action + " " + hour + ":" + minute} color="white" /> : <AddButton onClick={() => setPage("Select service")}/>} 
                    </AreaBox>

                    <AreaBox sx={{ flexDirection: "row", gap: 2, backgroundColor: "grey", width: "50vw", padding: "2%", borderRadius: 5, mb: 2 }}>
                        <AreaTypography variant="h3" text={action ? "Then" : "Then that"} color="white" />
                        {reaction ? <AreaTypography text={reaction} color="white" /> : action && <AddButton onClick={() => setPage("Select reaction")}/>} 
                    </AreaBox>

                    { action && reaction && <AreaButton text="Create area" onClick={() => createArea(hour, minute, channel, message)}/> }

                </AreaBox> }

                { page === "Select service" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Select a service" sx={{ mb: 5 }} />

                    {services && services.map<JSX.Element>((value: string) => {
                        return (<ServiceButton text={value} backgroundColor="#000" onClick={() => setPage("Select action")}/>);
                    })}

                </AreaBox> }

                { page === "Select action" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Select an action" sx={{ mb: 5 }} />

                    {actions && actions.map<JSX.Element>((value: string) => {
                        return (<ServiceButton text={value} backgroundColor="#000" onClick={() => setPage("Select hour")}/>);
                    })}

                </AreaBox> }

                { page === "Select hour" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Every day at" sx={{ mb: 5 }} />

                    <AreaTextField label="Hour" onChange={(s) => setHour(Number(s.target.value))} sx={{ width: "20vw" }} />
                    <AreaTextField label="Minute" onChange={(s) => setMinute(Number(s.target.value))} sx={{ width: "20vw" }} />

                    <AreaButton text="Valid" onClick={() => {setAction("Every day at"); setPage("Create"); }}/>

                </AreaBox> }

                { page === "Select reaction" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Select an action" sx={{ mb: 5 }} />

                    {reactions && reactions.map<JSX.Element>((value: string) => {
                        return (<ServiceButton text={value} backgroundColor="#000" onClick={() => setPage("Select channel")}/>);
                    })}

                </AreaBox> }

                { page === "Select channel" &&

                <AreaBox sx={{ height: "80vh", width: "98vw", gap: 2, mt: 10 }}>

                    <AreaTypography variant="h2" text="Select channel (id) and message" sx={{ mb: 5 }} />

                    <AreaTextField label="Channel (id)" onChange={(s) => setChannel(s.target.value)} sx={{ width: "20vw" }} />
                    <AreaTextField label="Message" onChange={(s) => setMessage(s.target.value)} sx={{ width: "20vw" }} />

                    <AreaButton text="Valid" onClick={() => {setReaction("Send a message on channel"); setPage("Create"); }}/>

                </AreaBox> }

                <Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Create;