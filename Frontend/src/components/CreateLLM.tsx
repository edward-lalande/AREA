import React, { useState } from "react";
import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";

import Footer from "./Footer";
import Navbar from "./Navbar";
import { AreaTextField } from "./elements/AreaTextFiled";
import { AreaButton } from "./elements/AreaButton";

import { useCookies } from "react-cookie";
import axios from "axios";

const CreateLLM: React.FC = () => {

    const [input, setInput] = useState<string>("");
    const [response, setResponse] = useState<string>("");
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const [cookies, setCookie] = useCookies();

    const sendRequest = () => {

        if (!input.trim()) {
            return;
        }

        setIsLoading(true);
        setResponse("");

        const url: string = "http://127.0.0.1:8080/areas-llm";

        const data = {
            user_token: cookies["token"],
            request: input
        };

        axios.post(url, data).then((res) => {
            if (res.status == 200) {
                setResponse("I have created your area!");
            } else {
                setResponse("Error, I can't create your area...");
            }
        });

        setIsLoading(false);
    }

    return (
    
        <AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }}>

            <AreaPaper sx={{ height: "96vh", width: "98vw" }}>

                <Navbar />

                <AreaBox sx={{ height: "96vh", width: "98vw", gap: 2 }}>

                <AreaTypography text="What AREA would you like to create?" sx={{ fontSize: "20px", fontWeight: "bold" }} />

                <AreaTextField
                    placeholder="Enter your request here..."
                    value={input}
                    onChange={(e) => {setInput(e.target.value)}}
                    sx={{ width: "40%" }}
                />

                <AreaButton text={isLoading ? "Processing..." : "Send"} onClick={sendRequest} disabled={isLoading} />

                <AreaBox sx={{ marginTop: "16px", textAlign: "center" }}>

                    {isLoading && <AreaTypography text="Chatbot is speaking..." />}

                    {!isLoading && response && (
                        <AreaTypography text={response} sx={{ color: response.includes("Error") ? "red" : "green" }} />
                    )}

                </AreaBox>

                </AreaBox>

                <Footer />

            </AreaPaper>

        </AreaBox>
    );
};

export default CreateLLM;
