import React from 'react';

export default class Reading extends React.Component<{}, { error: any, isLoaded: boolean, reading: any }> {
    constructor(props: any) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            reading: {},
        };
    }

    componentDidMount() {
        this.pollApi()
        setInterval(() => this.pollApi(), 10000);
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

    cToF(temp: number) {
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
                <div>
                    <span style={{ marginLeft: '.5rem', marginRight: '.5rem' }} >Humidity: {reading['Humidity']}</span>
                    <span style={{ marginLeft: '.5rem', marginRight: '.5rem' }} >Temperature: {this.cToF(reading['Temperature'])}F</span>
                </div>
            );
        }
    }
}
