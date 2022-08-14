package xavier.ricardo.paulobet.adapters;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import java.util.List;

import xavier.ricardo.paulobet.R;
import xavier.ricardo.paulobet.ScoresActivity;
import xavier.ricardo.paulobet.model.ScoreBoard;

public class ScoresAdapter extends BaseAdapter {
    private ScoresActivity context;
    private List<ScoreBoard> scores;

    public ScoresAdapter(ScoresActivity context, List<ScoreBoard> scores) {
        this.context = context;
        this.scores = scores;
    }

    @Override
    public int getCount() {
        return scores.size();
    }

    @Override
    public Object getItem(int i) {
        return scores.get(i);
    }

    @Override
    public long getItemId(int i) {
        return 0;
    }

    @Override
    public View getView(int i, View view, ViewGroup viewGroup) {
        if ((scores == null) || (scores.size() <= i)) {
            return null;
        }
        LayoutInflater inflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View v = inflater.inflate(R.layout.scoreboard, null);
        ScoreBoard scoreBoard = scores.get(i);
        TextView tvUser = v.findViewById(R.id.tvUser);
        tvUser.setText(scoreBoard.getUserId());
        TextView tvScore = v.findViewById(R.id.tvScore);
        tvScore.setText(scoreBoard.getScore().toString());
        return v;
    }
}
