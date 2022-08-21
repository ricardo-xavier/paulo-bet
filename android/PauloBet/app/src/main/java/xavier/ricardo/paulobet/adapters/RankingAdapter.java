package xavier.ricardo.paulobet.adapters;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import java.util.List;

import xavier.ricardo.paulobet.R;
import xavier.ricardo.paulobet.model.Ranking;

public class RankingAdapter extends BaseAdapter {
    private Context context;
    private List<Ranking> ranking;

    public RankingAdapter(Context context, List<Ranking> ranking) {
        this.context = context;
        this.ranking = ranking;
    }

    @Override
    public int getCount() {
        return ranking.size();
    }

    @Override
    public Object getItem(int i) {
        return ranking.get(i);
    }

    @Override
    public long getItemId(int i) {
        return 0;
    }

    @Override
    public View getView(int i, View view, ViewGroup viewGroup) {
        if ((ranking == null) || (ranking.size() <= i)) {
            return null;
        }
        LayoutInflater inflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View v = inflater.inflate(R.layout.scoreboard, null);
        Ranking r = ranking.get(i);
        TextView tvUser = v.findViewById(R.id.tvUser);
        tvUser.setText(r.getUserId());
        TextView tvScore = v.findViewById(R.id.tvScore);
        tvScore.setText(r.getScore().toString());
        return v;
    }
}
