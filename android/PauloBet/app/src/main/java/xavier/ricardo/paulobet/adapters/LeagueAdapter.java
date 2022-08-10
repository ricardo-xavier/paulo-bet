package xavier.ricardo.paulobet.adapters;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;

import java.util.List;

import xavier.ricardo.paulobet.SelectLeagueActivity;
import xavier.ricardo.paulobet.model.League;

import xavier.ricardo.paulobet.R;

public class LeagueAdapter extends BaseAdapter {
    private SelectLeagueActivity context;
    private List<League> leagues;

    public LeagueAdapter(SelectLeagueActivity context, List<League> leagues) {
        this.context = context;
        this.leagues = leagues;
    }

    @Override
    public int getCount() {
        return leagues.size();
    }

    @Override
    public Object getItem(int i) {
        return leagues.get(i);
    }

    @Override
    public long getItemId(int i) {
        return 0;
    }

    @Override
    public View getView(int i, View view, ViewGroup viewGroup) {
        if ((leagues == null) || (leagues.size() <= i)) {
            return null;
        }
        LayoutInflater inflater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        View v = inflater.inflate(R.layout.league, null);
        League league = leagues.get(i);
        TextView tvLeague = (TextView) v.findViewById(R.id.tvLeague);
        tvLeague.setText(league.getLeagueId());
        return v;
    }
}
