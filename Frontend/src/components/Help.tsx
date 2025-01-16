import { AreaBox } from "./elements/AreaBox";
import { AreaPaper } from "./elements/AreaPaper";
import { AreaTypography } from "./elements/AreaTypography";

import Footer from "./Footer";
import Navbar from "./Navbar";

const Help: React.FC = () => {

	return (

		<AreaBox width="100vw" height="100vh" sx={{ background: "#F2F2F2" }} >

			<AreaPaper sx={{ height: "96vh", width: "98vw" }}>

				<Navbar />

				<AreaBox sx={{ height: "80vh", width: "98vw", gap: 2 }}>

					<AreaTypography variant="h2" text="How to use Area?" sx={{ mb: 2 }} />

					<AreaBox sx={{ height: "60vh", width: "98vw", flexDirection: "row", gap: 5 }}>

						<AreaBox sx={{ height: "60vh", width: "40vw", backgroundColor: "#F2F2F2", border: "2px solid black", borderRadius: 5 }}>
							<AreaTypography variant="h4" text="How to create an area?" />

							<AreaTypography variant="h5" text="To create an area first you have to click on create button." />
							<AreaTypography variant="h5" text="Click on add button on (If) section to select an action." />
							<AreaTypography variant="h5" text="Then you have to choice a service and an action." />
							<AreaTypography variant="h5" text="Finaly you have to entry parameters for this action." />
							<AreaTypography variant="h5" text="You can do the same for the reaction on (Then) section." />
						
						</AreaBox>

						<AreaBox sx={{ height: "60vh", width: "40vw", backgroundColor: "#F2F2F2", border: "2px solid black", borderRadius: 5 }}>

							<AreaTypography variant="h4" text="How to manage areas?" />

							<AreaTypography variant="h5" text="To manage your areas you have to click on areas button." />
							<AreaTypography variant="h5" text="On this page you can see all your areas." />
							<AreaTypography variant="h5" text="Then you can manage them." />
							<AreaTypography variant="h5" text="You can delete one if you click on the trash button." />
						
						</AreaBox>
					
					</AreaBox>

				</AreaBox> 

				<Footer />

			</AreaPaper>

		</AreaBox>

	);

};

export default Help;
