import React from 'react';
import { LineChart, XAxis, YAxis, CartesianGrid, Line } from 'recharts';

export default class Linechart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            reading: {},
        };
    }

    componentDidMount() {
        this.pollApi()
    }

    pollApi() {
        fetch("http://ratatoskr:3000/api/readings")
            .then(res => res.json())
            .then(
                (result) => {
                    var newResults = []
                    for (var i = 0; i < result.length; i += 5) {
                        var newResult = {"Temperature": this.cToF(result[i]['Temperature'])}
                        newResults.push(newResult);
                    }
                    this.setState({
                        isLoaded: true,
                        reading: newResults
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            )
    }

    cToF(temp) {
        return temp * 9 / 5 + 32;
    }

    render() {
        const { error, isLoaded, reading } = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
            return (
                <LineChart width={500} height={300} data={reading}>
                    <XAxis dataKey="name" />
                    <YAxis domain={[50, 90]} />
                    <CartesianGrid stroke="#eee" strokeDasharray="5 5" />
                    <Line type="monotone" dataKey="Temperature" stroke="#8884d8" />
                </LineChart>
            );
        }
    }
}
