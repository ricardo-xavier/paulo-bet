package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.LeagueAdapter;
import xavier.ricardo.paulobet.model.GetLeaguesResponse;
import xavier.ricardo.paulobet.model.League;
import xavier.ricardo.paulobet.tasks.GetLeaguesTask;

public class SelectLeagueActivity extends AppCompatActivity {

    private SelectLeagueActivity context;
    private String user;
    private String token;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        context = this;
        setContentView(R.layout.activity_select_league);
        Intent intent = getIntent();
        user = intent.getStringExtra("user");
        token = intent.getStringExtra("token");

        ListView lvLeagues = findViewById(R.id.lvLeagues);
        lvLeagues.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                League league = (League) adapterView.getAdapter().getItem(i);
                Intent intent = new Intent(context, RankingActivity.class);
                intent.putExtra("league", league.getLeagueId());
                intent.putExtra("user", user);
                intent.putExtra("token", token);
                startActivity(intent);
                System.out.println(i);
            }
        });

        new GetLeaguesTask(this, user, token).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response == null) {
            Toast.makeText(this, "Sem resposta so servidor. Tente novamente mais tarde.", Toast.LENGTH_LONG).show();
            return;
        }
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            GetLeaguesResponse getLeaguesResponse = gson.fromJson(body, GetLeaguesResponse.class);
            ListView lvLeagues = findViewById(R.id.lvLeagues);
            LeagueAdapter adapter = new LeagueAdapter(this, getLeaguesResponse.getLeagues());
            lvLeagues.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }

    public void changePassword(View v) {
        Intent intent = new Intent(context, ChangePasswordActivity.class);
        intent.putExtra("user", user);
        intent.putExtra("token", token);
        startActivity(intent);
    }
}