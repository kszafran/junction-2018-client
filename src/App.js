import React, { Component } from "react";
import Header from "./components/Header";
import DeviceList from "./components/DeviceList";
import DeviceDetails from "./components/DeviceDetails";
import { Route } from "react-router-dom";
import "./App.css";
import Loader from "./components/Loader";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { devices: [], loader: true };
  }

  fetchJSON = async url => {
    const response = await fetch(url);
    try {
        return await response.json();
    } catch (err) {
        console.log(err);
        return undefined
    }
  };

  componentDidMount = async () => {
    const topology = await this.fetchJSON("https://cors-anywhere.herokuapp.com/https://peaceful-shelf-99858.herokuapp.com/topology");

    const clients = topology.nodes
        .filter(n => n.nodeType === "HOST" && n.additionalInfo && n.additionalInfo.macAddress)
        .map(n => n.additionalInfo.macAddress)
        .map(mac => "https://cors-anywhere.herokuapp.com/https://peaceful-shelf-99858.herokuapp.com/client-health/" + mac);

    const MOCK_DATA_URL =
      "https://gist.githubusercontent.com/ashleynguci/5fd6c84358844f9ac50f713b039bad8f/raw/0e1b2b19f0c79488b22d5326078e2b5206f2c48b/mock.json";

    const urlsArray = [MOCK_DATA_URL].concat(clients);
    const promiseArray = urlsArray.map(url => this.fetchJSON(url));
    const dataArray = await Promise.all(promiseArray);

    const devices = [...dataArray[0], ...dataArray.slice(1)].filter(d => d !== undefined);

    this.setState({
      devices: devices,
      loader: false
    });
  };

  getDeviceFromName = name => {
    return this.state.devices.find(e => e.name === name);
  };

  render() {
    const { devices, loader } = this.state;

    const Main = (
      <div className="App">
        <Header />
        <div className="wrapper">
          <Route
            exact
            path="/"
            render={() => <DeviceList devices={devices} />}
          />
          <Route
            path="/device/:name"
            render={props => (
              <DeviceDetails
                {...props}
                getDeviceFromName={this.getDeviceFromName}
              />
            )}
          />
        </div>
      </div>
    );

    return loader ? <Loader /> : Main;
  }
}

export default App;
