package xavier.ricardo.paulobet.model;

import java.util.List;

public class GetBetsResponse {
    private List<ScoreBoard> bets;

    public List<ScoreBoard> getBets() {
        return bets;
    }

    public void setBets(List<ScoreBoard> leagues) {
        this.bets = bets;
    }
}
