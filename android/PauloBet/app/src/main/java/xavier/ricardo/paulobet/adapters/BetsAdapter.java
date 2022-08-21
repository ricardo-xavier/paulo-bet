package xavier.ricardo.paulobet.adapters;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.ImageButton;
import android.widget.TextView;

import java.util.List;

import xavier.ricardo.paulobet.R;
import xavier.ricardo.paulobet.model.ScoreBoard;

public class BetsAdapter extends BaseAdapter {
    private Context context;
    private List<ScoreBoard> bets;

    public BetsAdapter(Context context, List<ScoreBoard> bets) {
        this.context = context;
        this.bets = bets;
    }

    @Override
    public int getCount() {
        return bets.size();
    }

    @Override
    public Object getItem(int i) {
        return bets.get(i);
    }

    @Override
    public long getItemId(int i) {
        return 0;
    }

    @Override
    public View getView(int i, View view, ViewGroup viewGroup) {
        if ((bets == null) || (bets.size() <= i)) {
            return null;
        }
        LayoutInflater inflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View v = inflater.inflate(R.layout.bet, null);
        ScoreBoard scoreBoard = bets.get(i);
        ImageButton ibEdit = v.findViewById(R.id.ibEdit);
        if (Boolean.FALSE.equals(scoreBoard.getEditable())) {
            ibEdit.setVisibility(View.INVISIBLE);
        }
        ibEdit.setTag(scoreBoard);
        TextView tvDate = v.findViewById(R.id.tvDate);
        tvDate.setText(scoreBoard.getDate());
        TextView tvMatch = v.findViewById(R.id.tvMatch);
        tvMatch.setText(scoreBoard.getMatchId());
        TextView tvBet = v.findViewById(R.id.tvBet);
        tvBet.setText(scoreBoard.getHome() + " x " + scoreBoard.getVisitors());
        TextView tvScore = v.findViewById(R.id.tvScore);
        tvScore.setText(scoreBoard.getScore().toString());
        return v;
    }
}
