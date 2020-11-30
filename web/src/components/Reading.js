import React from 'react';

export default class Reading extends React.Component {
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
        this.timer = setInterval(()=> this.pollApi(), 10000);
    }

    pollApi() {
        fetch("http://ratatoskr:3000/api/reading")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        reading: result
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
        console.log("testtest")
        console.log(error)
        console.log(isLoaded)
        console.log(reading)
        console.log("testtest")
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
            return (
                <ul>
                    <li>{reading['Humidity']}</li>
                    <li>{this.cToF(reading['Temperature'])}</li>
                </ul>
            );
        }
    }
}
