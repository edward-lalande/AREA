import { createTheme } from "@mui/material/styles";

const defaultTheme = createTheme({

	palette: {
		primary: {
			main: "#FF5500"
		},
		background: {
			default: "#FFFFFF",
			paper: "#F7F7F7"
		},
		text: {
			primary: "#000000",
			secondary: "#666666"
		},
	},
	typography: {
		fontFamily: "Avenir Next"
	},
	components: {
		MuiTextField: {
			styleOverrides: {
				root: {
					"& .MuiOutlinedInput-root": {
						borderRadius: "15px",
						"&.Mui-focused .MuiOutlinedInput-notchedOutline": {
							borderColor: "black"
						}
					},
					"& .MuiInputLabel-root": {
						color: "black",
						"&.Mui-focused": {
							color: "black"
						}
					},
				},
			},
		}
	}

});

export default defaultTheme;
