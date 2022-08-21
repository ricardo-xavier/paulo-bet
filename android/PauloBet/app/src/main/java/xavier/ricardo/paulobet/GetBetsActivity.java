package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.BetsAdapter;
import xavier.ricardo.paulobet.model.GetBetsResponse;
import xavier.ricardo.paulobet.model.ScoreBoard;
import xavier.ricardo.paulobet.tasks.GetBetsTask;

public class GetBetsActivity extends AppCompatActivity {
    private String league;
    private String user;
    private String login;
    private String token;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_get_bets);

        Intent intent = getIntent();
        league = intent.getStringExtra("league");
        user = intent.getStringExtra("user");
        login = intent.getStringExtra("login");
        token = intent.getStringExtra("token");

        new GetBetsTask(this, league, user, login, token).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response == null) {
            Toast.makeText(this, "Sem resposta so servidor. Tente novamente mais tarde.", Toast.LENGTH_LONG).show();
            return;
        }
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            GetBetsResponse betsResponse = gson.fromJson(body, GetBetsResponse.class);
            ListView lvBets = findViewById(R.id.lvBets);
            BetsAdapter adapter = new BetsAdapter(this, betsResponse.getBets());
            lvBets.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }

    public void bet(View v) {
        Intent intent = new Intent(this, BetActivity.class);
        ScoreBoard scoreBoard = (ScoreBoard) v.getTag();
        intent.putExtra("token", token);
        intent.putExtra("leagueId", league);
        intent.putExtra("userId", login);
        intent.putExtra("homeId", scoreBoard.getMatchId().split("-")[0]);
        intent.putExtra("visitorsId", scoreBoard.getMatchId().split("-")[1]);
        intent.putExtra("home", scoreBoard.getHome());
        intent.putExtra("visitors", scoreBoard.getVisitors());
        startActivityForResult(intent, 1);
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        new GetBetsTask(this, league, user, login, token).execute();
    }
}